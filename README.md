# Agro-Shield
Idoma centenary hackaton project


# FarmConnect — Crowd-Sourced Price Submission Service

Go backend module implementing crowd-sourced market price collection: farmers/field
agents submit prices they observe, submissions are automatically validated against
recent local history and submitter trust, and accepted prices roll up into daily
aggregates that power the "current price" / trend features.

## Run it

```bash
go run ./cmd/api
# Server listens on :8080
```

Submit a price:
```bash
curl -X POST localhost:8080/api/v1/prices/report \
  -d '{"commodity_id":"maize","market_id":"makurdi-central","submitter_id":"farmer-1","price":18000,"currency":"NGN","unit":"bag_100kg"}'
```

Get today's aggregate:
```bash
curl "localhost:8080/api/v1/prices/current?commodity_id=maize&market_id=makurdi-central"
```

### Admin: verify a submitter

Marks a field agent/vetted farmer as trusted, so their cold-start submissions
auto-accept instead of sitting `pending`:

```bash
curl -X POST localhost:8080/api/v1/admin/submitters/agent-1/verify \
  -d '{"trust_score":0.9}'
```

### Admin: review a pending or flagged report

```bash
curl -X POST localhost:8080/api/v1/admin/reports/<report_id>/review \
  -d '{"decision":"accepted","reason":"confirmed by phone with the reporting agent"}'
```

`decision` must be `"accepted"` or `"rejected"`. Only reports currently
`pending` or `flagged` can be reviewed.

> **No auth yet.** Both admin endpoints are wide open right now — they exist
> to prove out the moderation logic, not to go anywhere near production
> without an admin-only auth check in front of them first.

### Full end-to-end flow

```bash
# 1. An unverified farmer's first submission (no local history) -> pending
curl -X POST localhost:8080/api/v1/prices/report \
  -d '{"commodity_id":"maize","market_id":"makurdi-central","submitter_id":"farmer-1","price":18000,"currency":"NGN","unit":"bag_100kg"}'

# 2. Verify a field agent as trusted
curl -X POST localhost:8080/api/v1/admin/submitters/agent-1/verify -d '{"trust_score":0.9}'

# 3. That agent's submission now auto-accepts even with no local history yet
curl -X POST localhost:8080/api/v1/prices/report \
  -d '{"commodity_id":"maize","market_id":"makurdi-central","submitter_id":"agent-1","price":18000,"currency":"NGN","unit":"bag_100kg"}'

# 4. Current price now shows an aggregate
curl "localhost:8080/api/v1/prices/current?commodity_id=maize&market_id=makurdi-central"

# 5. Manually approve the farmer's earlier pending report from step 1
curl -X POST localhost:8080/api/v1/admin/reports/<report_id_from_step_1>/review \
  -d '{"decision":"accepted","reason":"matches agent-reported price"}'
```

Run tests:
```bash
go test ./... -v
```

## How the validation logic works (`internal/service/price_service.go`)

Every submission gets one of four statuses:

- **accepted** — within normal range of recent local prices, counted into aggregates
- **pending** — not enough local history yet to judge it (cold start), held for manual review
- **flagged** — a real outlier compared to recent local prices, held for review
- **rejected** — an extreme outlier (≥90% off local median), rejected outright regardless of trust

**Cold start problem**: a brand-new commodity/market pair has no price history to compare
against. Rather than blocking all early submissions, a *trusted, verified* submitter
(trust score ≥ 0.75) is auto-accepted even with zero history — this is how you'd seed
data through verified field agents before farmer-submitted volume builds up. Everyone
else's cold-start submissions go to `pending` for manual review.

**Outlier detection**: once ≥5 accepted prices exist for a commodity/market in the last
14 days, new submissions are compared against the **median** (not mean — a few bad
submissions shouldn't skew the reference point) of that recent history. The allowed
deviation band scales with the submitter's trust score, so a proven-reliable field
agent gets more leeway than a brand-new, unverified account before triggering review.

**Trust score**: every submitter has a 0–1 trust score (starts at 0.5, neutral).
It nudges up slightly on each accepted report and down on flagged/rejected ones,
so reliability compounds over time — a submitter who's consistently accurate needs
progressively less manual review.

## What's deliberately not built yet

- **Admin/moderation endpoints** — no way yet to manually verify a submitter or
  approve/reject a `pending`/`flagged` report via the API. Right now that only
  happens by calling the repository directly (see tests). This is the natural next
  piece: `POST /api/v1/admin/submitters/:id/verify`, `POST /api/v1/admin/reports/:id/review`.
- **Postgres implementation** — `domain.PriceRepository` / `domain.SubmitterRepository`
  are interfaces; only an in-memory implementation exists (`internal/repository/memory`).
  Swapping in Postgres means writing one new package that satisfies the same
  interfaces — the service and API layers don't change.
- **Scheduled aggregate recompute** — aggregates currently only recompute
  synchronously right after an accepted submission. A nightly job recomputing
  rolling 7-/30-day trend aggregates (for the trend-insight/sell-timing features)
  is the next layer on top of this.
- **Trend/comparison/recommendation endpoints** — this module covers ingestion
  and current price only, per today's scope.

## Structure

```
cmd/api/main.go                        entrypoint, wiring
internal/domain/                       entities + repository interfaces (no framework deps)
internal/repository/memory/            in-memory repository implementations
internal/service/price_service.go      submission validation, outlier detection, trust scoring
internal/service/stats.go              median / MAD helpers
internal/api/                          HTTP handlers + router
```
