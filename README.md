# limitlistener

This library makes it easy to limit the number of live threads for a given listener. This
can be useful in a number of cases to ensure not too many connections will be processed
at the same time rather than letting go create new goroutines like there is no tomorrow.
