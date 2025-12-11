import gleam/uri.{parse}
import gleam/option.{Some, None}
import gleam/string

pub fn normalize_url(url: String) -> Result(String, Nil) {
  case parse(url) {
    Ok(parsed_url) -> format_url(parsed_url)
    Error(err) -> Error(err)
  }
}

fn format_url(uri: uri.Uri) -> Result(String, Nil) {
  case uri.host {
    Some(host) -> Ok(trim_ending_slash(host <> uri.path) |> string.lowercase)
    None -> Error(Nil)
  }
}

fn trim_ending_slash(url: String) -> String {
  case string.ends_with(url, "/") {
    True -> string.slice(url, 0, string.length(url) - 1)
    False -> url
  }
}
