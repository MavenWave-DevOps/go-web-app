<!DOCTYPE html>
<html lang="en">

<head>
    <title>Maven Wave ToDo List v0.0.5-dev</title>
    <!-- JavaScript Bundle with Popper -->
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-1BmE4kWBq78iYhFldvKuhfTAU6auU8tT94WrHftjDbrCEXSU1oBoqyl2QvZ6jIW3" crossorigin="anonymous">    <!-- CSS only -->
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-ka7Sk0Gln4gmtz2MlQnikT1wXgYsOg+OMhuP+IlRH9sENBO0LRn5q+8nbTov4+1p" crossorigin="anonymous"></script>
</head>

<body>
<h1>Maven Wave ToDo's - Version 0.0.5-dev</h1>

<table class="table table-dark">
    <thead>
    <tr>
        <th>Title</th>
        <th>Description</th>
        <th>Status</th>
    </tr>
    </thead>
    <tbody id="tableBody">
    <td>Build ToDo App</td>
    <td>This will be an application we can test for Build, Release and Deployments!</td>
    <td>
        <select class="form-select" aria-label="Incomplete">
            <option value="volvo">Incomplete</option>
            <option value="saab">Complete</option>
        </select>
    </td>
    </tr>
    </tbody>
</table>
<br/>

<form class="form-control" action="/todo" method="post">
    <div class="form-group">
        <h3>Add a ToDo</h3>
        <label>Title:</label>
        <input class="form-control" type="text" name="title">
        <label>Description:</label>
        <input class="form-control" type="text" name="description">
        <button class="btn btn-primary" type="submit">Submit</button>
    </div>
</form>

<!--
  Finally, the last section is the script that will
  run on each page load to fetch the list of todos
  and add them to our existing table
 -->
<script>
    todoTable = document.querySelector("tbody")

    fetch("/todo")
        .then(response => response.json())
        .then(todoList => {
            //Once we fetch the list, we iterate over it
            todoList.forEach(todo => {
                // var currentTable = document.getElementById("tableBody")
                // Create the table row
                row = document.createElement("tr")

                // Create the table data elements for the title and
                // description columns
                title = document.createElement("td")
                title.innerHTML = todo.title
                description = document.createElement("td")
                description.innerHTML = todo.description

                // Add the data elements to the row
                row.appendChild(title)
                row.appendChild(description)
                // Finally, add the row element to the table itself
                todoTable.appendChild(row)
            })
        })
</script>

<footer>
    <p>Powered by ArgoCD</p>
    <img src={{.BucketUrl}} alt="Argo logo">
</footer>

</body>