
import normalize_url.{normalize_url}
import gleam/list
import gleam/io

pub fn normalize_url_test() {
  let tests = [
    #("remove scheme", "https://www.google.com", "www.google.com"),
    #("preserve path", "https://blog.boot.dev/path", "blog.boot.dev/path"),
    #("remove trailing slash", "https://blog.boot.dev/path/", "blog.boot.dev/path"),
    #("lowercase url", "https://BLOG.boot.dev/path/", "blog.boot.dev/path"),
  ]

  list.map(tests, fn(test_case) {
    io.println(test_case.0)
    let assert Ok(url) = normalize_url(test_case.1)
    assert url == test_case.2
  })
}
