# Articles


- [Exponential Backoff And Jitter](https://aws.amazon.com/tw/blogs/architecture/exponential-backoff-and-jitter/)
    + `Full Jitter`
        ```
            sleep = random_between(0, min(base * 2 ** attempt))
        ```
    + `Equal Jitter`
        ```
            tmp = min(cap, base * 2 ** attempt)
            sleep = tmp/2 + min(0, tmp/2)
        ```
    + `Decorrelated Jitter`
        ```
            sleep = min(cap, random_between(base, sleep * 3))
        ```

- [Multiple Region Multiple VPC Connectivity](https://aws.amazon.com/tw/answers/networking/aws-multiple-region-multi-vpc-connectivity/)