// All about config object

const AWS = require('aws-sdk');

AWS.config.apiVersions = {
    dynamodb: '2011-12-05',
    ec2: '2013-02-01',
    redshift: 'latest' // Best Practices for Managing AWS Access Keys
};



// Access key ID , Secret access key
// Get from IAM console -> Users -> choose IAM username -> choose Security credentials -> create access key

let credentials = new AWS.SharedIniFileCredentials({ profile: 'work-account' });
AWS.config.credentials = credentials; // global setting???

let myCredentials = new AWS.CognitoIdentityCredentials({ IdentityPoolId: 'IDENTITY_POOL_ID' });

// Global setting??
let config = new AWS.Config({
    credentials: myCredentials,
    region: 'us-west-2',
    maxRetries: 2,
    // logger:
    //     update:
});

// update current setting
config.update({ region: 'us-east-1' });
