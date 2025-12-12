import html_parser


pub fn get_first_paragraph(html: String) -> String {
  html |> html_parser.as_list |> find_first_paragraph
}

fn find_first_paragraph(in: List(html_parser.Element)) -> String {
  case in {
    [] -> ""
    [
      html_parser.StartElement("p", _, _),
      html_parser.Content(content),
      ..
    ] -> content
    [_, ..tail] -> find_first_paragraph(tail)
  }
}


pub fn get_h1_from_html(html: String) -> String {
  html |> html_parser.as_list |> find_h1
}

fn find_h1(in: List(html_parser.Element)) -> String {
  case in {
    [] -> ""
    [
      html_parser.StartElement("h1", _, _),
      html_parser.Content(content),
      ..
    ] -> content
    [_, ..tail] -> find_h1(tail)
  }
}
