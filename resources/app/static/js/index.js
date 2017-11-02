var index = {
    init: function() {

        document.getElementById("mailSenderList").value = ""
        document.getElementById("mailRecipientList").value = ""
        document.getElementById("mailTemplateList").value = ""

        // Wait for astilectron to be ready
        document.addEventListener('astilectron-ready', function() {
            // Listen
            index.listen();

            // Refresh list
            index.refreshList();

        })
    },

    //SENDING MAIL
    sendEmail: function() {
        var message = {
            "Sender" : document.getElementById("mailSenderField").value,
            "Recipient" : document.getElementById("mailRecipientField").value,
            "Subject" : document.getElementById("mailSubjectField").value,
            "Text" : document.getElementById("mailTextField").value,
            "ServerAddress" : document.getElementById("serverAddressField").value,
            "ServerPort" : document.getElementById("serverPortField").value,
        }
        astilectron.send({
            "name": "send.mail",
            "payload": message
        })
    },

    //MAIL SENDER
    setMailSenderListCustom: function() {
        var list = document.getElementById("mailSenderList");
        list.selectedIndex = list.length-1;
    },    
    setMailSenderField: function() {
        var list = document.getElementById("mailSenderList");
        document.getElementById("mailSenderField").value = list.options[list.selectedIndex].value;
    },

    //MAIL RECIPIENT
    setMailRecipientListCustom: function() {
        var list = document.getElementById("mailRecipientList");
        list.selectedIndex = list.length-1;
    },  
    setMailRecipientField: function() {
        var list = document.getElementById("mailRecipientList");
        document.getElementById("mailRecipientField").value = list.options[list.selectedIndex].value;
    },

    //MAIL TEXT TEMPLATE
    setMailTemplate: function() {
        var list = document.getElementById("mailTemplateList");
        document.getElementById("mailSubjectField").value = MAIL_TEMPLATES[list.selectedIndex][0];
        document.getElementById("mailTextField").value = MAIL_TEMPLATES[list.selectedIndex][1];
    },

    //MAIL SERVER
    setServerListCustom: function() {
        var list = document.getElementById("serverList");
        list.selectedIndex = list.length-1;
    },  
    setServerAddress: function() {
        var list = document.getElementById("serverList");
        document.getElementById("serverAddressField").value = MAIL_SERVERS[list.selectedIndex][0];
        document.getElementById("serverPortField").value = MAIL_SERVERS[list.selectedIndex][1];
    },

    listen: function() {
        astilectron.listen(function(message) {
            switch (message.name) {
                case "set.emailSender":
                    document.getElementById("list").innerHTML = message.payload;
                    break;
                case "sending.success":
                    document.getElementById("warningFade").style.display = "block";
                    document.getElementById("warningContent").textContent = message.payload;
                    break
                case "sending.error":
                    document.getElementById("warningFade").style.display = "block";
                    document.getElementById("warningContent").innerHTML = "<b align='left'>" + message.payload[0] + "</b>" + "</br></br>" + message.payload[1];
                    break
            }
        });
    },
    listenSetStyle: function(message) {
        document.body.className = message.payload;
    }
};

const MAIL_TEMPLATES = [
    ["...", "hi\njust looking on pieces of <your-name> and its fresh. you could be interested. -<name-of-curator>"],
    ["open for reccommendation?", "Hello,\nI thought that maybe you could be interested in a small tip of mine,\nif so check this portfolio out ;)\n\n<your-portfolio-link-here>\n\nCheers!\n<name-of-curator>"],
    ["just a tip", "hi \n\nthis got my eye yesterday,\ncheck it out: <your-portfolio-link-here>\n\n--<name-of-curator>"],
    ["A tip for an artist", "Hello,\n\ndo you mind a tip for an artist? I think you could be interested: \n\n<your-portfolio>\n\nWishes,\n<name-of-curator>\n<institution>"],
    ["artplay", "Hi,\n\nsending you an artist <your-portfolio-here>\n\nit's your turn ;)\n\n--<name-of-curator>"],
    ["", ""]
]

const MAIL_SERVERS = [
    ["mail.upcmail.cz", "25"],
    ["smtp.t-email.cz", "25"],
    ["", "25"]
]