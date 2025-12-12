
import get_urls_from_html.{get_urls_from_html, get_images_from_html}

pub fn get_urls_from_html_basic_test() {
  let input = "<html><body><a href=\"https://example.com\">Example</a></body></html>"
  let urls = get_urls_from_html(input)
  assert urls == ["example.com"]
}

pub fn get_urls_from_html_lots_of_urls_test() {
  let input = "<html><body><a href=\"https://example.com\">Example</a><a href=\"https://bar.com\">Example</a><a href=\"https://foo.com\">Example</a></body></html>"
  let urls = get_urls_from_html(input)
  assert urls == ["example.com", "bar.com", "foo.com"]
}

pub fn get_images_from_html_basic_test() {
  let input = "<html><body><img src=\"https://example.com/image.png\"></body></html>"
  let urls = get_images_from_html(input)
  assert urls == ["example.com/image.png"]
}
