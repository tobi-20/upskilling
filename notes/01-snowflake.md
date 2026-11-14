Why doesn't a normal database auto-increment ID work at Twitter's scale?
because they needed the capability of generating tens of thousands of IDs per second, unique IDs per second. I mean, um, and the conventional MySQL could not at the speed they needed for availability.

What are the 3 parts of a Snowflake ID, and why is each one there?
Three parts of a snowflake ID. You have the you have the timestamp, you have the worker number, and you have the sequence number.Timestamp is for tracing time in order to organize your tweets in the timeline with people that tweeted at the around the same time. Sequence number is used for thread to sort out the threads. Then worker number is also used for sorting, for giving unique IDs to the tweets, because each worker, when you use Zookeeper for worker, it automatically gives it a unique ID, a unique number, or a unique name.

What happens if the server's clock goes backwards?
I don't think it's going to do much harm because we still have the worker number and the sequence number. It's just that um tweets are the same timestamps. If the time goes backwards, the tweet might appear at the same time with other tweets that were initially at the right time. Oh,

What ID does Database A give its first tweet? What does Database B give its first tweet? What's the problem?
1, same 1. The two ids can clash

Follow-up 2: What's the name for that kind of weakness? ection 5.)
Lack of avilability

One worker gets 3 tweets in the same millisecond. Timestamp is the same for all 3. Worker ID is the same for all 3.
Follow-up 3: So what makes the 3 IDs different from each other? And when the clock moves to the next millisecond, what do you think happens to the sequence number?
I am not sure actually, when the clock moves to the next millisecond they all increment
1000, invalid, invalid
1000, out of sync, out of sync
