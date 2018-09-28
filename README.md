# A simple Go Web Crawler
> In progress...

running:
time go run main.go getInfoFromURL.go getListOfLinks.go sortPapers.go printPapers.go


#0 -     Company: CVRD ON N1
         Market Value: 317121000000.000000
#1 -     Company: PETROBRAS ON
         Market Value: 304719000000.000000
#2 -     Company: AMBEV S/A ON NM
         Market Value: 294004000000.000000
#3 -     Company: ITAUUNIBANCO PN N1
         Market Value: 283143000000.000000
#4 -     Company: AMBEV ON
         Market Value: 272509000000.000000
#5 -     Company: AMBEV PN
         Market Value: 271883000000.000000
#6 -     Company: PETROBRAS PN
         Market Value: 263368000000.000000
#7 -     Company: ITAUUNIBANCO ON N1
         Market Value: 247456000000.000000
#8 -     Company: BANCO BRADESCO S.A. PN N1
         Market Value: 191986000000.000000
#9 -     Company: CVRD PNA N1
         Market Value: 176290000000.000000

Time without any concurrency (go routines) used
real    4m33,679s
user    0m0,000s
sys     0m0,015s


dividing the slice of urls and running them using go routines:
With 5 go routines
real    1m14,511s
user    0m0,000s
sys     0m0,015s

with 10 go routines
real    1m18,038s
user    0m0,000s
sys     0m0,030s

with 20 go routines
real    1m17,083s
user    0m0,015s
sys     0m0,000s




explanation: 
A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished.

https://nathanleclaire.com/blog/2014/02/15/how-to-wait-for-all-goroutines-to-finish-executing-before-continuing/


Maximizing throughput is about getting rid of bottlenecks. First of all find where is most time is lost.

Sometimes running too many goroutines doesn’t make things faster but makes them slower, because goroutines start to compete for resources. Usually the fastest configuration is one that uses resources in the most effective way: if things are limited by computation then it's best to set the number of running goroutines equal to the number of CPU cores. But if it is limited by IO then choose the optimal number by measuring its performance.