# goccwc
Go solution to [Coding Challenges](https://codingchallenges.fyi/challenges/intro) first challenge: [build your own wc tool](https://codingchallenges.fyi/challenges/challenge-wc)

## Testing

### Step 1

```bash
goccwc % wc -c test.txt
  342190 test.txt
goccwc % go run ./cmd/ccwc -c test.txt
342190  test.txt
```

### Step 2

```bash
goccwc % wc -l test.txt
    7145 test.txt
goccwc % go run ./cmd/ccwc -l test.txt
7145    test.txt
```