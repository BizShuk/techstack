# Auto-Scalling

warm up period


[Auto Scalling limitation](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-account-limits.html)
### Cool Down period
Default of Cool Down period is 300 sec. This works for Dynamic Scalling.
takes effect starting when the last instance launches.

affected by lifecycle hook 

### Dynamic Scalling

### Scheduled Scalling


### termination policy 
Default: most number of instances in each availibility zone -> oldest launching configuration -> close to next billing hour -> random

### instance protection
`Don't protect from`:
- Manually termination throught the Amazon EC2 console, the erminate-instances command, or the TerminateInstances action.
-  Failed health check
-  spot instance interruption


### Lifecycle Hook
use it for scale out, pending status, and scale in, termination status.
Both of them have WAIT(The instance is paused until either you continue or the timeout period ends. One hour waiting time.) and PROCEED status. Check Lifecycle Hook status graph.

We can perform customized action:
- Define a CloudWatch event to invoke Lambda function
- Define a notification target for lifecycle hook.
- Create a script that runs on the instance as the instance starts. 

