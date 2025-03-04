/*
Copyright 2023 The Radius Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

module.exports = async ({ github, context }) => {
    if (context.eventName === 'issue_comment' && context.payload.action === 'created') {
        try {
            await handleIssueCommentCreate({ github, context });
        } catch (error) {
            console.log(`[handleIssueCommentCreate] unexpected error: ${error}`);
        }
    }
}

// Handle issue comment create event.
async function handleIssueCommentCreate({ github, context }) {
    const payload = context.payload;
    const issue = context.issue;
    const isFromPulls = !!payload.issue.pull_request;
    const commentBody = payload.comment.body;
    const username = context.actor;

    if (!commentBody) {
        console.log('[handleIssueCommentCreate] comment body not found, exiting.');
        return;
    }

    const commandParts = commentBody.split(/\s+/);
    const command = commandParts.shift();

    switch (command) {
        case '/assign':
            await cmdAssign(github, issue, isFromPulls, username);
            break;
        case '/ok-to-test':
            await cmdOkToTest(github, issue, isFromPulls, username);
            break;
        default:
            console.log(`[handleIssueCommentCreate] command ${command} not found, exiting.`);
            break;
    }
}

/**
 * Assign issue to the user who commented.
 * @param {*} github GitHub object reference
 * @param {*} issue GitHub issue object
 * @param {*} isFromPulls is the workflow triggered by a pull request?
 * @param {*} username is the user who trigger the command
 */
async function cmdAssign(github, issue, isFromPulls, username) {
    if (isFromPulls) {
        console.log('[cmdAssign] pull requests not supported, skipping command execution.');
        return;
    } else if (issue.assignees && issue.assignees.length !== 0) {
        console.log('[cmdAssign] issue already has assignees, skipping command execution.');
        return;
    }

    await github.rest.issues.addAssignees({
        owner: issue.owner,
        repo: issue.repo,
        issue_number: issue.number,
        assignees: [username],
    });
}

/**
 * Trigger e2e test for the pull request.
 * @param {*} github GitHub object reference
 * @param {*} issue GitHub issue object
 * @param {boolean} isFromPulls is the workflow triggered by a pull request?
 * @param {string} username is the user who trigger the command
 */
async function cmdOkToTest(github, issue, isFromPulls, username) {
    if (!isFromPulls) {
        console.log('[cmdOkToTest] only pull requests supported, skipping command execution.');
        return;
    }

    // Check if the user has permission to trigger e2e test with an issue comment
    const org = 'radius-project';
    console.log(`Checking team membership for: ${username}`);
    const isMember = await checkTeamMembership(github, org, process.env.TEAM_SLUG, username);
    if (!isMember) {
        console.log(`${username} is not a member of the ${teamSlug} team.`);
        return;
    }

    // Get pull request
    const pull = await github.rest.pulls.get({
        owner: issue.owner,
        repo: issue.repo,
        pull_number: issue.number,
    });

    if (pull && pull.data) {
        // Get commit id and repo from pull head
        const testPayload = {
            pull_head_ref: pull.data.head.sha,
            pull_head_repo: pull.data.head.repo.full_name,
            command: 'ok-to-test',
            issue: issue,
        };

        console.log('Creating repository dispatch event for e2e test');

        // Fire repository_dispatch event to trigger e2e test
        await github.rest.repos.createDispatchEvent({
            owner: issue.owner,
            repo: issue.repo,
            event_type: 'functional-tests',
            client_payload: testPayload,
        });

        console.log(`[cmdOkToTest] triggered E2E test for ${JSON.stringify(testPayload)}`);
    }
}

async function checkTeamMembership(github, org, teamSlug, username) {
    try {
        const response = await github.rest.teams.getMembershipForUserInOrg({
            org: org,
            team_slug: teamSlug,
            username: username,
        });
        return response.data.state === 'active';
    } catch (error) {
        console.log(`error: ${error}`)
        return false;
    }
}
