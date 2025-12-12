
import get_content.{get_h1_from_html, get_first_paragraph}

pub fn get_h1_from_html_basic_test() {
  let input = "<html><body><h1>Test Title</h1></body></html>"
  let h1 = get_h1_from_html(input)
  assert h1 == "Test Title"
}

pub fn get_h1_from_html_empty_test() {
  let input = "<html><body><p>No title</p></body></html>"
  let h1 = get_h1_from_html(input)
  assert h1 == ""
}

pub fn get_first_paragraph_basic_test() {
  let input = "<html><body><p>Test paragraph</p></body></html>"
  let paragraph = get_first_paragraph(input)
  assert paragraph == "Test paragraph"
}

pub fn get_first_paragraph_two_p_tags_test() {
  let input = "<html><body><p>Test paragraph</p><div>foo</div><p>Second paragraph</p></body></html>"
  let paragraph = get_first_paragraph(input)
  assert paragraph == "Test paragraph"
}
