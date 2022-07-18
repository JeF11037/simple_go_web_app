function erase()
{
    var tr = document.getElementById("adding");
    for (let td of tr.children)
    {
        td.children[0].value = "";
    }
}

function rowOnFocus(button, trId)
{
    button.setAttribute("hidden", "hidden");
    button.nextElementSibling.removeAttribute("hidden");
    for (let td of document.querySelector('[data-user-id="'+trId+'"]').children)
    {
        var input = td.children[0];
        if (input.nodeName == "INPUT")
            input.removeAttribute("readonly");
        if (input.nodeName == "SELECT")
            input.removeAttribute("disabled");  
    }
}

function validateName(input)
{
    var regex = /^[a-zA-Z]{1,100}$/;
    if(regex.test(input.value))
        input.style.borderColor = "green";
    else
        input.style.borderColor = "red";
}

function validateEmail(input)
{
    var regex = /^(([^<>()[\]\\.,;:\s@']+(\.[^<>()[\]\\.,;:\s@']+)*)|('.+'))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    if(regex.test(input.value))
        input.style.borderColor = "green";
    else
        input.style.borderColor = "red";
}

function translatePage()
{
    var language = localStorage.getItem('language');
    if (language == null)
    {
        localStorage.setItem("language", "en")
        translatePage();
    }
    else
    {
        document.getElementById("language").value = language;
        $.getJSON('translation/translation.json', function(data) {
            $('[data-translation]').each(function(i, obj) {
                var translation = data;
                $.each($(obj).data("translation").split(" "), function(j, path) {
                    translation = translation[path]
                });
                if ($(obj).children().length > 0)
                    $(obj).children("a").text(translation[language]);
                else
                    $(obj).text(translation[language]);
            });
        });
    }
}

function changeLanguage(value)
{
    localStorage.setItem("language", value);
    translatePage();
}