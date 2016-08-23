package main

const uploadPage = `
<!DOCTYPE html>
<html>
   <head>
      <title>Upload File Page</title>
   </head>
   <body>
         <form action="/aws/url" method="post" enctype="multipart/form-data">
         <div><label for="aws_file">Upload your File: </label><input type="file" name="aws_file" size="40"></div>
         <br></br>
         <div><label for="submit"><input type="submit" value="Submit"></div>
      </form>
   </body>
</html>
`
