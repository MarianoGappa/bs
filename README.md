# bs - byte summer

Reads off STDIN, looks for human-readable file sizes, sums them up (rounding) and outputs the sum to STDOUT in -h form
```
~/Code/go/src/github.com/MarianoGappa/bs $ ll -h
total 5832
drwxr-xr-x   8 marianol  staff   272B  8 May 21:02 .
drwxr-xr-x   9 marianol  staff   306B  8 May 20:14 ..
drwxr-xr-x  14 marianol  staff   476B  8 May 21:03 .git
drwxr-xr-x  11 marianol  staff   374B  8 May 20:59 .idea
-rw-r--r--   1 marianol  staff   1.1K  8 May 20:14 LICENSE
-rwxr-xr-x   1 marianol  staff   2.8M  8 May 21:02 bs
-rw-r--r--   1 marianol  staff   409B  8 May 20:15 bs.iml
-rw-r--r--   1 marianol  staff   1.3K  8 May 20:59 main.go
~/Code/go/src/github.com/MarianoGappa/bs $ ll -h | bs
2.8M
```

# bs is not du

* Does not recurse subdirectories
* Allows you to pipe-filter STDIN
```
~/Code/go/src/github.com/MarianoGappa/bs master* $ du -hs
4.0M	.
~/Code/go/src/github.com/MarianoGappa/bs master* $ ll -h | bs
2.8M
~/Code/go/src/github.com/MarianoGappa/bs master* $ ll -h | grep \.go | bs
1.3K
```
* Note that most linux distros don't use the `B` suffix for sizes in bytes (i.e. less than 1K), so on Linux you might get slightly smaller numbers.

# Installing
```
go get github.com/MarianoGappa/bs
```
