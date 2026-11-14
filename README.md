# Upskilling: Rebuilding Real-World Engineering Incidents

I'm rebuilding, from scratch, the systems behind famous engineering incidents and scaling problems, such as Twitter's ID generation, Cloudflare's regex outage and GitHub's 43-second network partition.

For each build I study what happened and why, implement a small version of the system, test the edge cases that caused the real problem, and write down the trade-offs and known limitations.

## Repository structure

```
.
├── projects/              # one folder per build, each a standalone Go module
│   └── 01-snowflake/      # Twitter Snowflake: distributed unique ID generator
│       ├── README.md      # design decisions, bit layout, known limitations
│       ├── snowflake.go
│       └── snowflake_test.go
└── notes/                 # my study notes and Q&A for each build
    └── 01-snowflake.md
```

- **`projects/NN-name/`:** the code. Each one has its own `README.md` explaining the incident, the design and what's deliberately left out.
- **`notes/NN-name.md`:** the questions I worked through before and while building, answered in my own words.

## Roadmap

Ordered from easiest to hardest. Each build carries forward concepts from the earlier ones.

| # | Build | Based on | Core concept | Status |
|---|---|---|---|---|
| 1 | [Snowflake ID generator](projects/01-snowflake) | Twitter | Bit manipulation, distributed IDs | ✅ Done |
| 2 | Rate limiter (token bucket + sliding window) | Google / Shopify | Algorithms, Redis | ⏳ Next |
| 3 | Circuit breaker + flaky service | Netflix | Resilience patterns, state machines | |
| 4 | Cache stampede demo + singleflight | Facebook 2010 | Concurrency, thundering herd | |
| 5 | Flash sale (fix overselling) | Shopify | Transactions, locking, race conditions | |
| 6 | Feature flags + staged rollout | Knight Capital / CrowdStrike | Safe deploys | |
| 7 | Backup + automated restore test | GitLab | Ops, Postgres, replication | |
| 8 | Content-addressed store / mini-git | Dropbox | Hashing, storage | |
| 9 | Regex engine (Thompson NFA) | Cloudflare | Automata, complexity | |
| 10 | Shard router | Instagram / Figma | Sharding, routing | |
| 11 | Mini DNS server | Facebook 2021 | Networking, binary protocols | |
| 12 | Heartbleed in C + fuzzing | OpenSSL | Memory safety, security | |
| 13 | LSM-tree KV store | Discord | Storage engines | |
| 14 | Raft | GitHub 2018 | Distributed consensus | |

## How each build is taken further

Each project can be pushed through four levels:

1. **Works:** a minimal correct implementation.
2. **Measured:** tests, benchmarks and load tests, so I know the numbers.
3. **Broken on purpose:** reproduce the original failure and show the fix surviving it.


## Running the code

Each project is its own Go module. Use Go 1.27 or later.

```bash
cd projects/01-snowflake
go test ./...
```

## References

- Martin Kleppmann, *Designing Data-Intensive Applications*
- [danluu/post-mortems](https://github.com/danluu/post-mortems): a curated list of public postmortems
