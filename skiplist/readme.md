
## SKIP List



## Test

```sh
=== RUN   TestSkipList_Delete
=== RUN   TestSkipList_Delete/delete_from_empty_list
before delete: -9223372036854775808:HEAD->9223372036854775807:NIL
before delete: -9223372036854775808:HEAD->9223372036854775807:NIL
=== RUN   TestSkipList_Delete/delete_-_key_not_exists
before delete: -9223372036854775808:HEAD->0:test->6:A->8:B->9223372036854775807:NIL
before delete: -9223372036854775808:HEAD->0:test->6:A->8:B->9223372036854775807:NIL
=== RUN   TestSkipList_Delete/delete_-_delete_the_element
before delete: -9223372036854775808:HEAD->0:test->6:A->8:B->9223372036854775807:NIL
before delete: -9223372036854775808:HEAD->0:test->8:B->9223372036854775807:NIL
=== RUN   TestSkipList_Delete/delete_-_delete_the_element:0
before delete: -9223372036854775808:HEAD->0:test->6:A->8:B->9223372036854775807:NIL
before delete: -9223372036854775808:HEAD->6:A->8:B->9223372036854775807:NIL
=== RUN   TestSkipList_Delete/delete_-_delete_the_element:6
before delete: -9223372036854775808:HEAD->0:test->6:A->8:B->9223372036854775807:NIL
before delete: -9223372036854775808:HEAD->0:test->8:B->9223372036854775807:NIL
--- PASS: TestSkipList_Delete (0.00s)
    --- PASS: TestSkipList_Delete/delete_from_empty_list (0.00s)
    --- PASS: TestSkipList_Delete/delete_-_key_not_exists (0.00s)
    --- PASS: TestSkipList_Delete/delete_-_delete_the_element (0.00s)
    --- PASS: TestSkipList_Delete/delete_-_delete_the_element:0 (0.00s)
    --- PASS: TestSkipList_Delete/delete_-_delete_the_element:6 (0.00s)
```