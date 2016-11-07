# Resilient Services in Go

## Abstract

Resilience is not about never failing, but how do you recover from it.
How you can prevent your services from locking down or exhausting all
its resources ? How to perform graceful service degradation ?
Can this kind of behaviour be tested properly ?

On Go we have some new features, like Contexts, that helps us
to model timeouts and cancellation properly.

They can be combined with other useful features as select and channels
to model timeouts and resource pools, which can be essential
to provide proper service degradation instead of total failure of the system.

On this talk I try to answer this questions using new features available
on Go 1.7, direct from production ready software.


## Presenting
