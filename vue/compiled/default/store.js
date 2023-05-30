
    var api_domain = "/api/"

    function getData(key, json) {
        try {
            if (localStorage) {
                if (json) {
                    return JSON.parse(localStorage.getItem(key))
                } else {
                    return localStorage.getItem(key)
                }
            } else {
                return null
            }
        } catch (error) {
            return null
        }
    }

    function setItem(key, value) {
        try {
            localStorage.setItem(key, value)
        } catch (error) {
            console.log("Error while saving data on localstorage")
        }
    }

    function removeItem(key) {
        try {
            localStorage.removeItem(key)
        } catch (error) {
            console.log("Error while removing data on localstorage")
        }
    }

    function playSound(silent) {
        //  
        var audio = new Audio("/sound/notification.mp3");
        if (silent) {
            audio.volume = 0.001
        } else {
            audio.volume = 1
        }
        audio.play();
    }

    export default {
        state: {
            conn : null,
            connecting : false,
            user: getData('user', true),
            account: getData('account', true),
            pages: [],
            locations: [],
            notification_count: 0,
            notifications: {},
            sessiontypes: [],
            levels: [],
            sublevels: [],
            subjects: [],
            recent_posts: [],
            recent_events: [],
            recent_sessions: [],
            usersinchat : [],
            letest_message : false,
            last_message: "",
            last_message_error: false,
            location: getData('location', true),
            country: getData('country'),
            district: getData('district'),
            locality: getData('locality'),
            sublocality: getData('sublocality'),
            code: getData('code'),
            subject: "",
            sublevel: "",
            level: "",
            search: "",
            page_title: "Page",
            broadcast_channel: false,
            service_worker: null,
        },
        getters: {
            getAccountUrl: (state) => (account) => {
                var url = document.location.protocol + "//"
                if (account.Active || account.active) {
                    if (account.Domain) {
                        url += account.Domain
                    } else if (account.Subdomain) {
                        url += account.Subdomain + "." + window.application.BaseDomin
                    } else if (account.domain) {
                        url += account.domain
                    } else if (account.subdomain) {
                        url += account.subdomain + "." + window.application.BaseDomin
                    }
                } else if (account.AccountID) {
                    url += window.application.BaseDomin + "/account/" + account.AccountID
                } else if (account.ID) {
                    url += window.application.BaseDomin + "/account/" + account.ID
                }
                return url
            },
            getFBLink: (state) => (path) => {
                var url = document.location.protocol + "//" + document.location.host + "/" + path
                var link = "https://www.facebook.com/sharer/sharer.php?u=" + encodeURI(url)
                return link
            },
            getTwitterLink: (state) => (path) => {
                var url = document.location.protocol + "//" + document.location.host + "/" + path
                var link = "https://twitter.com/share?url=" + encodeURI(url)
                return link
            },
            getGooglePlusLink: (state) => (path) => {
                return ""
                var url = document.location.protocol + "//" + document.location.host + "/" + path
                var link = "https://www.facebook.com/sharer/sharer.php?u=" + encodeURI(url)
                return link
            },
            getPinterestLink: (state) => (path) => {
                return ""
                var url = document.location.protocol + "//" + document.location.host + "/" + path
                var link = "https://www.facebook.com/sharer/sharer.php?u=" + encodeURI(url)
                return link
            },
            getFullDateTime: (state) => (timestamp) => {
                var date
                if (typeof (timestamp) == 'string') {
                    date = (new Date(timestamp))
                } else if (timestamp == 0 || timestamp == undefined || timestamp == "") {
                    return "Unknown"
                } else {
                    date = (new Date(timestamp * 1000))
                }
                return months[date.getMonth()] + " " + date.getDate() + " " + date.getFullYear() + " - " + (date.getHours() % 12).toString() + ":" + date.getMinutes() + (((date.getHours() / 12) > 1) ? "PM" : "AM")
            },

            getDuration: (state) => (course) => {
                var value = course.EndsOn - course.StartFrom
                value = Math.ceil(value / 60 / 60 / 24)
                return value
            },
            getSessions: (state) => (course) => {
                var value = course.EndsOn - course.StartFrom
                value = Math.ceil(value / 60 / 60 / 24 / 7)
                return value
            },
            getSessionsCompleted: (state) => (course) => {
                var now = new Date().getTime() / 1000
                var value = now - course.StartFrom
                if (value < 0) {
                    return 0
                }
                value = Math.ceil(value / 60 / 60 / 24 / 7)
                return value
            },
            getSessionsRemaning: (state) => (course) => {
                var now = new Date().getTime() / 1000
                var value = course.EndsOn - now
                if (value < 0) {
                    return 0
                }
                value = Math.ceil(value / 60 / 60 / 24 / 7)
                return value
            },
            getTime: (state) => (timestamp) => {
                var date
                if (typeof (timestamp) == 'string') {
                    if (timestamp.length == 5) {
                        timestamp = timestamp.split(":")
                        var hr = parseInt(timestamp[0])
                        var sifix = "AM"
                        if (hr > 12) {
                            hr = hr - 12
                            sifix = "PM"
                        }
                        if (hr == 12) {
                            sifix = "noon"
                        }
                        if (hr == 24) {
                            hr = "00"
                            sifix = "midnight"
                        }
                        if (hr < 10) {
                            hr = "0" + hr
                        }
                        return hr + ":" + timestamp[1] + sifix
                    }
                    date = (new Date(timestamp))
                } else if (timestamp == 0 || timestamp == undefined || timestamp == "") {
                    return "Unknown"
                } else {
                    date = (new Date(timestamp * 1000))
                }
                return (date.getHours() % 12).toString() + ":" + date.getMinutes() + (((date.getHours() / 12) > 1) ? "PM" : "AM")
            },
            getFullDate: (state) => (timestamp) => {
                var date
                if (typeof (timestamp) == 'string') {
                    date = (new Date(timestamp))
                } else if (timestamp == 0 || timestamp == undefined || timestamp == "") {
                    return "Unknown"
                } else {
                    date = (new Date(timestamp * 1000))
                }
                return months[date.getMonth()] + " " + date.getDate() + " " + date.getFullYear()
            },
            getMonth: (state) => (timestamp) => {
                var monthNumber = 1
                if (typeof (timestamp) == 'string') {
                    monthNumber = (new Date(timestamp)).getMonth()
                } else if (timestamp == 0 || timestamp == undefined || timestamp == "") {
                    monthNumber = 1
                } else {
                    monthNumber = (new Date(timestamp * 1000)).getMonth()
                }
                return months[monthNumber]
            },
            getDate: (state) => (timestamp) => {
                if (typeof (timestamp) == 'string') {
                    return (new Date(timestamp)).getDate()
                } else if (timestamp == 0 || timestamp == undefined || timestamp == "") {
                    return 1
                } else {
                    return (new Date(timestamp * 1000)).getDate()
                }
            },
        },
        mutations: {
            reset_notifications(state) {
                var keys = Object.keys(state.notifications)
                keys.forEach(url => {
                    if(window.location.pathname == url){
                        delete(state.notifications[url])
                    }
                });
                state.notification_count = Object.keys(state.notifications).length
            },
            set_notification(state, payload) {
                if(state.notification[payload.notification.Url]){
                    state.notification[payload.notification.Url].Content  + "\n" + payload.notification.Content
                }else{
                    state.notifications[payload.notification.Url] = payload.notification
                }
                state.notification_count = Object.keys(state.notifications).length
            },
            set_chats(state, payload) {
                if(Array.isArray(payload)){
                    payload.forEach(user => {
                        var usersinchat = state.usersinchat.filter(userinchat => (userinchat.ID == user.ID));
                        if(!usersinchat.length){
                            state.usersinchat.push(user)
                        }else{
                            if (usersinchat.length > 1) {
                                alert("Check out there are more users then expected")
                            }
                            var userinchat = usersinchat[0]
                            var userinchatUpdatedAt = (new Date(userinchat.UpdatedAt)).getTime()
                            var userUpdatedAt = (new Date(user.UpdatedAt)).getTime()
                            if (userinchatUpdatedAt < userUpdatedAt) {
                                userinchat.Name = user.Name
                                userinchat.Photo = user.Photo
                                userinchat.Online = user.Online
                                userinchat.LastActiveOn = user.LastActiveOn

                                if(user.ChatID > 0){
                                    userinchat.Content = user.Content
                                    userinchat.Type = user.Type
                                    userinchat.RecivedAt = user.RecivedAt
                                    userinchat.SeenAt = user.SeenAt
                                    userinchat.ReplayTo = user.ReplayTo
                                    userinchat.UpdatedBy = user.UpdatedBy
                                    userinchat.CreatedBy = user.CreatedBy
                                    userinchat.CreatedFor = user.CreatedFor
                                    userinchat.MessageID = user.MessageID
                                    userinchat.TotalMessages = user.TotalMessages
                                    userinchat.TotalUsers = user.TotalUsers
                                    userinchat.CreatedAt = user.CreatedAt
                                    userinchat.UpdatedAt = user.UpdatedAt
                                    userinchat.DeletedAt = user.DeletedAt
                                }


                                userinchat.UpdatedAt = user.UpdatedAt
                                userinchat.CreatedAt = user.CreatedAt
                                userinchat.Message = user.Message
                                userinchat.Files = user.Files
                                userinchat.Sent = user.Sent
                                userinchat.SentAt = user.SentAt
                                userinchat.RecivedAt = user.RecivedAt
                                userinchat.SeenAt = user.SeenAt
                                userinchat.Total = user.Total
                            }
                        } 
                    });
                    state.usersinchat.sort((a, b) => { return a.CreatedAt - b.CreatedAt; });
                }
            },
            set_message(state, payload) {
                var chats = state.usersinchat.filter(userinchat => (userinchat.ChatID == payload.message.ChatID));
                if(chats.length){
                    var chat = chats[0]
                }else{
                    if(payload.sender){
                        chat = {
                            ID : payload.sender.ID,
                            Name: payload.sender.Name,
                            Photo : payload.sender.Photo,
                            Online : payload.sender.Online,
                            LastActiveOn : payload.sender.LastActiveOn,
                            IsSupportagent : payload.sender.IsSupportagent,
                            ChatID : payload.message.ChatID,
                            MessageID : 0,
                            TotalMessages : 0,
                            TotalUsers : 0,
                            Ispublic : false,
                        }
                    }
                }

                if(!chat.Messages){
                    chat.Messages = []
                }
                var item_index = chat.Messages.findIndex(message => (message.ID == payload.message.ID))
                if(item_index > -1){
                    chat.Messages.splice(item_index, 1);
                }
                chat.Messages.push(payload.message)
                
                if(chat.MessageID < payload.message.ID){
                    chat.MessageID = payload.message.ID
                    chat.Content = payload.message.Content
                    chat.Type = payload.message.Type
                    chat.RecivedAt = payload.message.RecivedAt
                    chat.SeenAt = payload.message.SeenAt
                    chat.ReplayTo = payload.message.ReplayTo
                    chat.UpdatedBy = payload.message.UpdatedBy
                    chat.CreatedBy = payload.message.CreatedBy
                    chat.CreatedFor = payload.message.CreatedFor
                    chat.CreatedAt = payload.message.CreatedAt
                    chat.UpdatedAt = payload.message.UpdatedAt
                    chat.DeletedAt = payload.message.DeletedAt
                    chat.TotalMessages++ 
                }
                state.letest_message = payload.message

                var url = "/support/"+ chat.ID
                if(window.location.pathname != url){
                    if(state.notifications[url]){
                        state.notifications[url].Count++
                        state.notifications[url].Title = chat.Name + " send "+ state.notifications[url].Count+" messages"
                        state.notifications[url].Content  += "\n" + payload.message.Content
                    }else{
                        state.notifications[url] = {
                            Url : url,
                            Image : chat.Photo,
                            Title : chat.Name + " sent you a message" ,
                            Content : payload.message.Content, 
                            Count : 1 
                        }
                    }
                    state.notifications = JSON.parse(JSON.stringify(state.notifications))
                    state.notification_count = Object.keys(state.notifications).length
                }
            },
            set_conn(state, payload) {
                if(payload == "connecting"){
                    state.connecting = true
                }else{
                    state.connecting = false
                    state.conn = payload
                    if(payload){
                        state.user.Online = true
                    }else{
                        state.user.Online = false
                    }
                }
            },
            set_page_title(state, payload) {
                state.page_title = payload
            },
            set_country(state, payload) {
                state.country = payload
                setItem('country', state.country)
            },
            set_district(state, payload) {
                state.district = payload
                setItem('district', state.district)
            },
            set_locality(state, payload) {
                state.locality = payload
                setItem('locality', state.locality)
            },
            set_sublocality(state, payload) {
                state.sublocality = payload
                setItem('sublocality', state.sublocality)
            },
            set_code(state, payload) {
                state.code = payload
            },
            set_subject(state, payload) {
                state.subject = payload
            },
            set_sublevel(state, payload) {
                state.sublevel = payload
            },
            set_level(state, payload) {
                state.level = payload
            },
            set_search(state, payload) {
                state.search = payload
            },
            set_recent_sessions(state, payload) {
                if (Array.isArray(payload)) {
                    state.recent_sessions = payload
                }
            },
            set_subjects(state, payload) {
                if (Array.isArray(payload)) {
                    state.subjects = payload
                    var sublevelmap = {}
                    var levelmap = {}
                    state.subjects.forEach(subject => {
                        if (sublevelmap[subject.SubLevelName]) {
                            sublevelmap[subject.SubLevelName].Subjects.push(subject)
                            sublevelmap[subject.SubLevelName].SessionCount += subject.SessionCount
                        } else {
                            sublevelmap[subject.SubLevelName] = {
                                ID: subject.SubLevelID,
                                LevelID: subject.LevelID,
                                Name: subject.SubLevelName,
                                LevelName: subject.LevelName,
                                SessionCount: subject.SessionCount,
                                Subjects: [subject]
                            }
                        }
                    });
                    sublevelmap = Object.values(sublevelmap)
                    sublevelmap.sort(function (a, b) { return b.SessionCount - a.SessionCount });
                    sublevelmap.forEach(sublevel => {
                        sublevel.Subjects.sort(function (a, b) { return b.SessionCount - a.SessionCount });
                    });
                    state.sublevels = sublevelmap
                    state.sublevels.forEach(sublevel => {
                        if (levelmap[sublevel.LevelName]) {
                            levelmap[sublevel.LevelName].SubLevels.push(sublevel)
                            levelmap[sublevel.LevelName].SubjectCount += sublevel.SubjectCount
                            levelmap[sublevel.LevelName].SessionCount += sublevel.SessionCount
                        } else {
                            levelmap[sublevel.LevelName] = {
                                ID: sublevel.LevelID,
                                Name: sublevel.LevelName,
                                SessionCount: sublevel.SessionCount,
                                SubjectCount: sublevel.SubjectCount,
                                SubLevels: [sublevel]
                            }
                        }
                    });
                    levelmap = Object.values(levelmap)
                    levelmap.sort(function (a, b) { return b.SessionCount - a.SessionCount });
                    state.levels = levelmap
                }
            },
            set_pages(state, payload) {
                if (Array.isArray(payload)) {
                    state.pages = payload
                }
            },
            set_sessiontypes(state, payload) {
                if (Array.isArray(payload)) {
                    state.sessiontypes = payload
                }
            },
            set_locations(state, payload) {
                if (Array.isArray(payload)) {
                    state.locations = payload
                }
            },
            set_location(state, payload) {
                state.location = payload
                if (_.isEmpty(state.location)) {
                    removeItem('location')
                } else {
                    setItem('location', JSON.stringify(state.location))
                }
            },
            set_recent_posts(state, payload) {
                if (Array.isArray(payload)) {
                    state.recent_posts = payload
                }
            },
            set_recent_events(state, payload) {
                if (Array.isArray(payload)) {
                    state.recent_events = payload
                }
            },
            set_current_user(state, payload) {
                state.user = payload
                if (_.isEmpty(state.user)) {
                    removeItem('user')
                } else {
                    setItem('user', JSON.stringify(state.user))
                }
            },
            set_current_account(state, payload) {
                state.account = payload
                if (_.isEmpty(state.account)) {
                    removeItem('account')
                } else {
                    setItem('account', JSON.stringify(state.account))
                }
            },
            set_last_message(state, payload) {
                if (!payload) {
                    state.last_message = ""
                    state.last_message_error = false
                } else if (payload.Message) {
                    state.last_message = payload.Message
                    if (payload.Status == 2) {
                        state.last_message_error = false
                    } else {
                        state.last_message_error = true
                    }
                }
            },
            clear_search(state) {

                state.country = ""
                state.district = ""
                state.locality = ""
                state.sublocality = ""
                state.code = ""
                state.level = ""
                state.sublevel = ""
                state.subject = ""
                state.search = ""

                removeItem('country')
                removeItem('district')
                removeItem('locality')
                removeItem('sublocality')
                removeItem('code')
                removeItem('level')
                removeItem('sublevel')
                removeItem('subject')
                removeItem('search')

            }
        },
        actions: {
            set_ws_responce(context, data) {
                var userid = 0
                if(data.closed){
                    context.commit('set_conn', false)
                    return
                }
                if (data.message) {
                    context.commit('set_message', data)
                }
                if(data.online){
                    var usersinchat = context.state.usersinchat.filter(userinchat => (userinchat.ID == data.online));
                    if(usersinchat.length ){
                        usersinchat[0].Online = true
                    }
                }
                if(data.offline){
                    var usersinchat = context.state.usersinchat.filter(userinchat => (userinchat.ID == data.offline));
                    if(usersinchat.length ){
                        usersinchat[0].Online = false
                    }
                }
                if (data.usersinchat && Array.isArray(data.usersinchat)) {
                    context.state.usersinchat = data.usersinchat
                    context.state.usersinchat.sort((a, b) => { return a.CreatedAt - b.CreatedAt; });
                }
            },
            init_web_socket: _.debounce((context) => {
                if(context.state.connecting || !context.state.user){
                    return
                }
                context.commit('set_conn', "connecting")
                if (!window["WebSocket"]) {
                    console.log("Your browser does not support Live chat.");
                    return
                }
                context.state.broadcast_channel = new BroadcastChannel('my_notification_channel');
                context.state.broadcast_channel.addEventListener('message', (data) => {
                    if (context.state.user && !context.state.conn) {

                        if(data == "opened"){
                            context.state.user.Online = true
                        }else if(data == "closed"){
                            context.state.user.Online = false
                            //context.dispatch('init_web_socket', data) 
                        }else if(data.data == "opened"){
                            context.state.user.Online = true
                        }else if(data.data == "closed"){
                            context.state.user.Online = false
                            //context.dispatch('init_web_socket', data) 
                        }else{
                            try {
                                console.log(data)
                                var data = JSON.parse(data)
                                context.dispatch('set_ws_responce', data) 
                            } catch (error) {
                                console.log(data)
                                console.log(error)
                            }
                        }
                    }
                    
                })
                window.onbeforeunload = function () {
                    if(context.state.conn){
                        context.state.broadcast_channel.postMessage("closed")
                    }
                };
                //context.state.service_worker = navigator.serviceWorker.register('service.js');
                var conn
                if (location.host.includes("localhost")) {
                    conn = new WebSocket("ws://" + document.location.host + "/ws/" + context.state.user.ID);
                } else {
                    conn = new WebSocket("wss://" + document.location.host + "/ws/" + context.state.user.ID);
                }
                conn.onopen = (event) => {
                    console.log("WebSocket is open now.");
                    context.state.broadcast_channel.postMessage("opened")
                    if (context.state.user) {
                        context.state.user.Online = true
                        //context.commit('set_current_user', context.state.user)
                        var data = {
                            api: "set_user",
                        }
                        data.token = context.state.user.Token
                        if (context.state.account && context.state.account.ID) {
                            data.account_id = context.state.account.ID
                        } else if (window.application.Account.ID) {
                            data.account_id = window.application.Account.ID
                        }
                        conn.send(JSON.stringify(data));
                        context.commit('set_conn', conn)
                    } else {
                        conn.close()
                    }
                };
                conn.onclose = (evt) => {
                    if (context.state.user) {
                        context.state.user.Online = true
                    }
                    context.state.broadcast_channel.postMessage("closed")
                    console.log("WebSocket Connection closed.");
                    context.commit('set_conn', false)
                };
                conn.onerror = (evt) => {
                    console.log("WebSocket Connection failed.");
                    console.log(evt)
                };
                conn.onmessage = (evt) => {
                    console.log(evt)
                    console.log("WebSocket Connection recived a message.");
                    if (evt.data == "who are you?") {
                        console.log("User not set tring to set again")
                        if (context.state.user) {
                            var data = {
                                api: "set_user",
                            }
                            data.token = context.state.user.Token
                            if (context.state.account && context.state.account.ID) {
                                data.account_id = context.state.account.ID
                            } else if (window.application.Account.ID) {
                                data.account_id = window.application.Account.ID
                            }
                            conn.send(JSON.stringify(data));
                        } else {
                            conn.close()
                        }
                    } else {
                        try {
                            context.state.broadcast_channel.postMessage(evt.data)
                            try {
                                var data = JSON.parse(evt.data)
                                context.dispatch('set_ws_responce', data)
                                if (data.message) {
                                    if (data.message.CreatedBy != context.state.user.ID) {
                                        playSound()
                                    } else {
                                        playSound(false)
                                    }
                                }
                            } catch (error) {
                                console.log(error)
                            }
                            
                        } catch (error) {
                            console.log(error)
                        }
                    }
                };
            }, 1000),
            logout(context) {
                return context.dispatch("call", {
                    api: "logout",
                    data: {}
                })
            },
            call(context, { api, data } = {}) {
                if (context.state.user) {
                    data.token = context.state.user.Token
                }

                if (context.state.account && context.state.account.ID) {
                    data.account_id = context.state.account.ID
                } else if (window.application.Account.ID) {
                    data.account_id = window.application.Account.ID
                }
                var url = api_domain + api

                console.log(url);
                return fetch(url, {
                    method: 'post',
                    headers: {
                        "Content-type": "application/x-www-form-urlencoded; charset=UTF-8"
                    },
                    body: JSON.stringify(data)
                }).then(response => response.json()).then(function (data) {
                    console.log(data);
                    if (data.User) {
                        context.commit('set_current_user', data.User)
                        if (!data.User.Online) {
                            context.dispatch('init_web_socket')
                        }
                    } else {
                        context.commit('set_current_user', null)
                    }
                    if (data.Account) {
                        context.commit('set_current_account', data.Account)
                    } else {
                        context.commit('set_current_account', null)
                    }
                    if (data.Message) {
                        context.commit('set_last_message', data)
                    }
                    return data
                });
            }
        },
    }
