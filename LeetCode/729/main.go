package main

import "fmt"

type MyCalendar struct {
    bookings [][]int
}

func Constructor() MyCalendar {
    return MyCalendar{[][]int{}}
}

func (this *MyCalendar) Book(start int, end int) bool {
    for _, booking := range this.bookings {
        if clash(start, end, booking[0], booking[1]) {
            return false
        }
    }
    booking := []int{start, end}
    this.bookings = append(this.bookings, booking)
    return true
}

func clash(start1, end1, start2, end2 int) bool {
    return start2 < end1 && start1 < end2
}

func main() {
    calendar := Constructor()
    bookings := [][]int{
        {10, 20},
        {15, 25},
        {20, 30},
    }

    for _, booking := range bookings {
        start, end := booking[0], booking[1]
        if calendar.Book(start, end) {
            fmt.Printf("Booking [%d, %d] successful.\n", start, end)
        } else {
            fmt.Printf("Booking [%d, %d] failed due to clash.\n", start, end)
        }
    }
}
