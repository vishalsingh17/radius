import radius as radius

param context object

var basename = context.resource.name

// This is not a realistic user scenario (creating a Radius resource in a recipe). We're
// doing things this way to test the UCP functionality without using cloud resources.
resource extender 'Applications.Link/extenders@2022-03-15-privatepreview' = {
  name: '${basename}-module'
  properties: {
    application: context.application.id
    environment: context.environment.id
    resourceProvisioning: 'manual'
    message: 'hello from module resource'
  }
}
