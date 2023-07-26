# tugas-api-go

### # Run
By default, this API runs on endpoint '/users' and on port '9000'<br>
You can access it with 'http://{url}:9000/users'

---
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
  * `nim` can't use number 0
  * `nama` can`t use empty string
  * `jurusan` can't use empty string
