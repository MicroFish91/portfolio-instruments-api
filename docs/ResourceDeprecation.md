# ❌ Resource Deprecation

Certain resources have an `is_deprecated` field which may be set if a user requests to delete a resource that is still being actively utilized. This is very common for benchmarks, 
accounts, and holdings - all of which are used in the producing of snapshots.  In such cases where a resource is still being used when a delete request has been issued, 
the resource will instead be marked deprecated.  In follow-up queries, deprecated resources will no longer show up except in the context for which they are still being
utilized.  Whenever the utilizing resource (e.g. a snapshot) is deleted, the logic will automatically check to see if any dependent resources can be removed automatically
(by checking if the dependent resources have the field `is_deprecated` set to true).  In such cases, the server will again attempt to delete these now dangling resources.