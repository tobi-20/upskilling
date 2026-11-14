# Snowflake ID Generator in Go

A simple implementation of Twitter's Snowflake: a way for many servers to create unique, time-sortable 64-bit IDs.

## Synopsis

In 2010 Twitter needed tens of thousands of unique IDs per second across many machines. One database handing out `1, 2, 3…` becomes a bottleneck and a single point of failure. With several databases, each one starts at `1` and their IDs collide. Snowflake lets every server generate IDs on its own.

## Bit layout

```
| 1 bit unused | 41 bits timestamp (ms) | 10 bits worker ID | 12 bits sequence |
```

- **Timestamp:** milliseconds since a custom epoch, so IDs sort by creation time.
- **Worker ID (0–1023):** unique per machine, so two machines never clash.
- **Sequence (0–4095):** counts IDs within the same millisecond on the same machine.

Two IDs can only be equal if all three parts match, and the generator makes sure that never happens.

The result is a **signed `int64`** with the top bit left at 0, as in the original. Java has no unsigned 64-bit type, and many databases don't either, so the ID stays positive everywhere.

## Design decisions

- **Custom epoch:** timestamps count from 2026-11-12, not 1970 as in Unix. Starting from 1970 would already use up ~56 of the ~69.7 years that 41 bits can hold. The epoch is a hardcoded constant and **must never change**: changing it breaks time-ordering and may produce duplicate IDs.
- **Fixed-width `uint64`:** Go's plain `uint` is only 32 bits on 32-bit machines, where the shifted timestamp would silently lose bits. The tests pass with `GOARCH=386`.
- **Unsigned inputs:** negative numbers in two's complement start with a run of 1s, which would flood the neighbouring fields when combined with `|`.
- **The worker ID is validated once, in `NewGenerator`:** a misconfigured server fails at startup instead of handing out bad IDs.

## Known limitations

- **It errors instead of waiting when a millisecond is full.** After 4096 IDs in one millisecond, `NextID` returns an error. Twitter's version spins until the next millisecond (`tilNextMillis`).
- **Worker IDs are assigned by hand.** Twitter used ZooKeeper to hand out worker IDs and to check that no two servers shared one.
- **A backwards clock is rejected, not tolerated.** If NTP steps the clock back, `NextID` returns an error rather than risk reissuing an ID.
- **No datacenter bits.** The original split the 10 bits into 5 for the datacenter and 5 for the worker.
- **The generator is not safe for concurrent use.** Calling it from several goroutines at once would need a mutex.

## Run the tests

```bash
go test ./...
GOARCH=386 go test ./...   # simulate a 32-bit machine
```

## References

- Twitter Engineering, "Announcing Snowflake" (2010)
- github.com/twitter-archive/snowflake: the original Scala `IdWorker`
