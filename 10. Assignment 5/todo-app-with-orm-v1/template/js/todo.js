const checkAuth = async () => {
  const response = await fetch("http://localhost:8080/user/session/valid", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: 'same-origin',
  });

  const myJson = await response.json();
  if (myJson.message != "Token Valid") {
    window.location.href = '/page/login';
  }
};

checkAuth()

const logoutAction = async () => {
  const response = await fetch("http://localhost:8080/user/logout", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: 'same-origin',
  });

  const myJson = await response.json();
  console.log(myJson)

  window.location.href = '/page/login';
};

const listTodoAction = async () => {
  const response = await fetch("http://localhost:8080/todo/list", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: 'same-origin',
  });

  const myJson = await response.json();
  
  document.getElementById("myUL").innerHTML = ''
  for (i = 0; i < myJson.length; i++) {
    var li = document.createElement("li");
    if (myJson[i].done) {
      li.className = "checked";
    } 
    li.appendChild(document.createTextNode(myJson[i].ID));
    li.appendChild(document.createTextNode(". "));
    li.appendChild(document.createTextNode(myJson[i].task));
    document.getElementById("myUL").appendChild(li);
    var span = document.createElement("SPAN");
    var txt = document.createTextNode("\u00D7");
    span.className = "close";
    span.appendChild(txt);
    li.appendChild(span);
  }

  closeGenerate()
};

listTodoAction()

const deleteTodoAction = async (id) => {
  const response = await fetch("http://localhost:8080/todo/remove?id=" + id, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: 'same-origin',
  });

  const myJson = await response.json();

  if (myJson.error === undefined) {
    pushNotify("success", myJson.username, myJson.message)
    listTodoAction()
  } else {
    pushNotify("error", "Error", myJson.error)
  }
};

function closeGenerate() {
  var closeElement = document.getElementsByClassName("close");
  for (var i = 0; i < closeElement.length; i++) {
    closeElement[i].onclick = function() {
      var div = this.parentElement;
      let contentParent = this.parentElement.innerHTML.split(".");
      deleteTodoAction(contentParent[0])
      div.style.display = "none";
    }
  }
};

const toggleDone = async (id = 0, toggle = false) => {
  const response = await fetch("http://localhost:8080/todo/change-status", {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: 'same-origin',
    body: JSON.stringify({
      id: Number(id),
      done: toggle,
    }),
  });

  const myJson = await response.json();
  if (myJson.error === undefined) {
    pushNotify("success", myJson.username, myJson.message)
  } else {
    pushNotify("error", "Error", myJson.error)
  }
};

// Add a "checked" symbol when clicking on a list item
var list = document.querySelector("ul");
list.addEventListener(
  "click",
  function (ev) {
    if (ev.target.tagName === "LI") {
      let content = ev.target.outerText.split(".");
      ev.target.classList.toggle("checked");
      let id = content[0]
      console.log(content[0], ev.target.className);
      if (ev.target.className == "checked") {
        toggleDone(id, true)
      } else {
        toggleDone(id, false)
      }
    }
  },
  false
);

const newElement = async (taskDesc = "") => {
  var inputValue = document.getElementById("myInput").value;
  const response = await fetch("http://localhost:8080/todo/add", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: 'same-origin',
    body: JSON.stringify({
      task: inputValue,
      done: false,
    }),
  });

  const myJson = await response.json();

  if (myJson.error === undefined) {
    pushNotify("success", myJson.username, myJson.message)
    listTodoAction()
  } else {
    pushNotify("error", "Error", myJson.error)
  }
};


function pushNotify(status, title, message) {
  new Notify({
    status: status,
    title: title,
    text: message,
    effect: 'fade',
    speed: 300,
    customClass: null,
    customIcon: null,
    showIcon: true,
    showCloseButton: true,
    autoclose: false,
    autotimeout: 3000,
    gap: 20,
    distance: 20,
    type: 1,
    position: 'right top'
  })
}