Go has a very strange way to initilize slices. You create the array, and then you initilize it into a slice with make(). From there, you can append and so on.

My question here: Why do we have to initilize the slice with make()?

Answer: https://share.google/aimode/eYZzoeOO3k8sMhdYm

That was really interesting to find out, Go is very strange in this way, but it's good to know that utlizing make is helpful in terms of memory and cpu usage.