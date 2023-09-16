# Cloud Formation


### wait condition
Maximum time: 12 hr
`CREATE_IN_PROGRESS`, after it receives the number of successful signals, it'll create further stacks.  Otherwise, it'll be `CREATE_FAILED`.

DependOn <resources>

ways to send signal:
1. curl
2. cfn-signal
3. SignalResource API


https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/deploying.applications.html

### troubleshoot 
1. diable stack rolling back by appending --disable-rollback
2. login and check logs
3. check log of /var/log/cfn-init.log
4. check log of /var/lib/cloud/instance/scripts/part-001
5. check correctness of metadata by using cfn-get-metadata