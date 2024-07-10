import React from 'react';
import { FileUpload } from 'primereact/fileupload';
import 'primereact/resources/themes/lara-light-indigo/theme.css';

const FileUploadComponent = () => {

    const onUpload = (e) => {
        const uploadedFiles = e.files;
        // Handle the uploaded files
        console.log('Uploaded Files:', uploadedFiles);
        // You can process the files here, e.g., send them to the server or update the state
    };

    return (
        <div className="card">
            <FileUpload name="demo[]" url="./upload" onUpload={onUpload} multiple accept="image/*" maxFileSize={1000000} />
        </div>
    );
}

export default FileUploadComponent;
