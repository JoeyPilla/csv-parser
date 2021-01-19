import React from "react";

import "./App.css";

function App() {
  return (
    <div>
      <form
        action="/receive_multiple"
        method="post"
        encType="multipart/form-data"
      >
        <label for="file">Filenames:</label>
        <input type="file" name="multiplefiles" id="multiplefiles" multiple />
        <input type="submit" name="submit" value="Submit" />
      </form>
    </div>
  );
}

export default App;
