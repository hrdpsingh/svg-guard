use quick_xml::events::Event;
use quick_xml::reader::Reader;
use quick_xml::writer::Writer;
use std::{fs, io::Cursor};

fn remove_script(svg: &str) -> String {
    let mut reader = Reader::from_str(svg);
    let mut writer = Writer::new(Cursor::new(Vec::new()));
    let mut buffer = Vec::new();
    let mut skip_content = false;
    
    reader.config_mut().trim_text(true);

    loop {
        match reader.read_event_into(&mut buffer) {
            Ok(Event::Start(ref e)) if e.name().as_ref() == b"script" => {
                skip_content = true;
            }
            Ok(Event::End(ref e)) if e.name().as_ref() == b"script" => {
                skip_content = false;
            }
            Ok(Event::Empty(ref e)) if e.name().as_ref() == b"script" => {}
            Ok(Event::Eof) => break,
            Ok(event) => {
                if !skip_content {
                    writer.write_event(event).expect("Error writing to output");
                }
            }
            Err(e) => panic!("Error at position {}: {:?}", reader.buffer_position(), e),
        }
        buffer.clear();
    }
    
    let result = writer.into_inner().into_inner();
    String::from_utf8(result).expect("Generated invalid UTF-8")
}

fn main() {    
    let svg = fs::read_to_string("test-cases/xss.svg").expect("Unable to read file");
    let clean_svg = remove_script(&svg);
    println!("Sanitized SVG:\n{}", clean_svg);
}