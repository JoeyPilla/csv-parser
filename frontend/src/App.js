import React, { useState } from "react";
import styled from "styled-components";

function App() {
  const [year, setYear] = useState("");
  const [files, setFiles] = useState([]);

  const handleFiles = (e) => {
    const files = e.target.files;
    setFiles(Array.from(files).map((file) => file.name));
  };

  const Files = files.map((file) => <File>{file}</File>);
  return (
    <Div>
      <Form
        action={`/receive_multiple?year=${year}`}
        method="post"
        encType="multipart/form-data"
      >
        <FileInput>
          <Label>
            <Input
              type="file"
              name="multiplefiles"
              id="multiplefiles"
              multiple
              onChange={(e) => handleFiles(e)}
            />
            Upload Files
          </Label>
          <Filelist>
            <FileHeader>Files to be Uploaded</FileHeader>
            {Files}
          </Filelist>
        </FileInput>
        <label>
          Year:
          <YearInput
            value={year}
            onChange={(e) => {
              const value = e.target.value;
              if (/^-?\d*$/.test(value)) {
                setYear(value);
              }
            }}
            placeholder="YYYY"
          />
        </label>
        {year.length > 0 && year.length !== 4 && (
          <Error>Year must be in YYYY format</Error>
        )}
        <SubmitInput
          type="submit"
          name="submit"
          value="Submit"
          disabled={year.length !== 4}
        />
      </Form>
    </Div>
  );
}

export default App;

const Div = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100vh;
  width: 100vw;
`;

const Input = styled.input`
  display: none;
  cursor: pointer;
`;

const YearInput = styled.input`
  margin-left: 8px;
  padding-left: 8px;
`;

const SubmitInput = styled.input`
  border: none;
  background: #404040;
  color: #ffffff !important;
  margin-top: 8px;
  padding: 12px;
  text-transform: uppercase;
  border-radius: 6px;
  display: inline-block;
  transition: all 0.3s ease 0s;
  &:hover:not([disabled]) {
    cursor: pointer;
    color: #404040 !important;
    font-weight: 700 !important;
    letter-spacing: 3px;
    background: none;
    -webkit-box-shadow: 0px 5px 40px -10px rgba(0, 0, 0, 0.57);
    -moz-box-shadow: 0px 5px 40px -10px rgba(0, 0, 0, 0.57);
    transition: all 0.3s ease 0s;
  }
`;

const Label = styled.label`
  width: 35%;
  border: none;
  background: #404040;
  color: #ffffff !important;
  margin-top: 8px;
  padding: 12px;
  text-transform: uppercase;
  text-align: center;
  border-radius: 6px;
  display: inline-block;
  transition: all 0.3s ease 0s;
  cursor: pointer;
  &:hover {
    color: #404040 !important;
    font-weight: 700 !important;
    letter-spacing: 3px;
    background: none;
    -webkit-box-shadow: 0px 5px 40px -10px rgba(0, 0, 0, 0.57);
    -moz-box-shadow: 0px 5px 40px -10px rgba(0, 0, 0, 0.57);
    transition: all 0.3s ease 0s;
  }
`;

const File = styled.p`
  padding-left: 8px;
  margin: 0px;
`;

const Error = styled.p`
  color: tomato;
  padding-left: 8px;
  margin: 0px;
`;

const FileHeader = styled.h3`
  text-decoration: underline;
  padding-left: 8px;
  margin: 0px;
  margin-bottom: 8px;
`;

const Filelist = styled.div`
  display: flex;
  width: 35%;
  min-width: 35%;
  text-overflow: ellipsis;
  overflow: hidden;
  flex-direction: column;
  flex-wrap: wrap;
  padding: 8px;
  min-height: 1px;
`;

const FileInput = styled.div`
  display: flex;
  width: 70%;
  padding-bottom: 16px;
  justify-content: center;
  align-items: center;
`;

const Form = styled.form`
  display: flex;
  align-items: center;
  flex-direction: column;
  width: 50%;
`;
