# tugas-api-go

### # Run
By default, this API runs on endpoint '/users' and on port '9000'<br>
You can access it with 'http://{url}:9000/users'

---
### # Add User
* `'POST'`
* `'http:{url}:9000/users'`
* `body request`
  <pre>
    {
      "nim": 2023,
      "nama": "sholikul", 
      "jurusan": "sistem informasi"
    }
  </pre>
  * `nim` can't be filled with an number 0
  * `nama` can't be filled with an empty string
  * `jurusan` can't be filled with an empty string
