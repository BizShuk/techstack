let aws = require('aws-sdk');




// EC2 sample
var ec2_regionA = new AWS.EC2({region: 'ap-southeast-2', maxRetries: 15, apiVersion: '2014-10-01'});
var ec2_regionB = new AWS.EC2({region: 'us-east-1', maxRetries: 15, apiVersion: '2014-10-01'});



let s3 = new aws.S3(s3.config);











