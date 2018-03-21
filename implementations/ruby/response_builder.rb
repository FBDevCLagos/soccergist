class ListTemplateBuilder
  attr_accessor :elements

  def initialize
    @elements = []
  end

  def add_new_element(title:, image_url:, subtitle:, btn: nil)
    element_builder = ElementBuilder.new(
      title: title,
      image_url: image_url,
      subtitle: subtitle
    )

    element_builder.add_btn(btn) if btn
    @elements << element_builder.build
  end

  def build
    {
      message: {
        attachment: {
          type: 'template',
          payload: {
            template_type: 'list',
            top_element_style: 'compact',
            elements: elements
          }
        }
      }
    }
  end
end

class ElementBuilder
  attr_accessor :buttons, :title, :image_url, :subtitle

  def initialize(title:, image_url:, subtitle:)
    @buttons = []
    @title = title
    @image_url = image_url
    @subtitle = subtitle
  end

  def add_btn(btn)
    @buttons << btn
  end

  def build
    {
      title: title,
      subtitle: subtitle,
      image_url: image_url,
      buttons: buttons
    }
  end
end

class ButtonResponseBuilder
  def self.build_postback(title:, postback:)
    {
      type: 'postback',
      title: title,
      payload: postback
    }
  end
end
