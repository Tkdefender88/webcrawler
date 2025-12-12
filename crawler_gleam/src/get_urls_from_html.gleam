
import html_parser
import gleam/list
import normalize_url.{normalize_url}

pub fn get_urls_from_html(html: String) -> List(String) {
  html |> html_parser.as_list |> find_urls([])
}

fn find_urls(in: List(html_parser.Element), urls: List(String)) -> List(String) {
  case in {
    [] -> urls
    [
      html_parser.StartElement("a", [html_parser.Attribute("href", url)], _),
      ..tail
    ] -> {
      case url |> normalize_url {
        Ok(normalized_url) -> find_urls(tail, list.append(urls, [normalized_url]))
        Error(_) -> find_urls(tail, urls)
      }
    }
    [_, ..tail] -> find_urls(tail, urls)
  }
}

pub fn get_images_from_html(html: String) -> List(String) {
  html |> html_parser.as_list |> find_images([])
}

fn find_images(in: List(html_parser.Element), urls: List(String)) -> List(String) {
  case in {
    [] -> urls
    [
      html_parser.StartElement("img", [html_parser.Attribute("src", url)], _),
      ..tail
    ] -> {
      case url |> normalize_url {
        Ok(normalized_url) -> find_images(
          tail, list.append(urls, [normalized_url]))
        Error(_) -> find_images(tail, urls)
      }
    }
    [_, ..tail] -> find_images(tail, urls)
  }
}
