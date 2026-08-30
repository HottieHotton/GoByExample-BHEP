When breaking down Mutexes, I had Gemini break it down, which helps explain the use case for it:
> Imagine you and a group of roommates are working on a single, shared whiteboard.
>
>If two of you try to write on or erase the same spot at the exact same time, you'll end up with scrambled, illegible scribbles. To prevent this, you might make a simple rule: Whoever is holding the marker is the only one allowed to touch the board. When they're done, they hand the marker to the next person.
>
>In Go, a Mutex (short for Mutual Exclusion) is that marker.

When you "lock", that person now has a marker, when it's "unlocked", it's being passed or not being used.

NOTES: Go has a race detector to ensure your code passes any race condition issues:
>Go Race Detector Integration: Go includes a built-in tool (go run -race main.go) that inspects your code at runtime. Using mutexes properly ensures your code passes race detection cleanly.

Gemini also outlined the following:
> #### A Quick Rule of Thumb for Go Beginners
>Go has a famous proverb: "Do not communicate by sharing memory; instead, share memory by communicating."
>
>- Use Channels when you want to pass data or signals back and forth between goroutines.
>
>- Use Mutexes when you have a shared data structure (like a cache, map, or counter) that multiple goroutines just need to read or update safely.