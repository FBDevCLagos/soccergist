# On code organisation!
_- Oscar_

As you work more, interacting with the Messenger platform, you eventually realise the repetitive and verbose nature of the response templates. Thus far, we have seen the text template, button template and this week, the list template. As the project grows, you'll find your codebase littered with these unsightly boilerplates and they start being harder to manage.

One good approach to arrest this is to have dedicated functions which wrap around the messenger api to help build out these responses. This way, you can easily reuse the logic and avoid duplicating your efforts whenever you want to use the same template. In our case (so far), you would have a `build_text_response` function, a `build_button_template_response` function, and a `build_list_template_response` function each of which takes the required arguments and return the desired response structure. As you work with more templates you will also find some similar response structure which you can also extract to its own function to nicely clean up your codebase. You can then move the builder functions to their own class or file (depending on your language's abstraction provision).

The end result is that you can then focus on your business logic without worrying or being distracted by implementation details.