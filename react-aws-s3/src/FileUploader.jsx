import React, { useState } from 'react';
import axios from 'axios';
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

const FileUploader = () => {
    const [file, setFile] = useState(null);
    const [presignedURL, setPresignedURL] = useState('');
    const [errorMessage, setErrorMessage] = useState('');
    const toastOptions = {
        position: "top-right",
        autoClose: 3000,
        pauseOnHover: true,
        draggable: true,
        theme: "dark",
    };

    const handleFileChange = (event) => {
        setFile(event.target.files[0]);
    };

    const handleUpload = async () => {
        if (!file) {
            setPresignedURL("")
            setErrorMessage('Please select a file to upload');
            return;
        }

        try {
            const response = await axios.post('http://localhost:9000/upload-file', {
                file_name: file.name,
                file_size: file.size,
            });

            if (response?.data?.meta?.code === 0) {
                toast.error(
                    response?.data?.meta?.message,
                    toastOptions
                );  
                return
            }

            await axios.put(response?.data?.data?.pre_signed_url, file, {
                headers: {
                    'Content-Type': file.type
                },
            });
            setPresignedURL(response?.data?.data?.pre_signed_url)
            setFile(null);
            setErrorMessage('');
        } catch (error) {
            setErrorMessage('Error uploading file: ' + error.message);
        }
    };

    return (
        <div className="container mx-auto mt-10 p-5 bg-gray-100 rounded-lg">
            <h1 className="text-2xl font-semibold mb-5">File Uploader</h1>
            <input type="file" onChange={handleFileChange} className="mb-3" />
            <button onClick={handleUpload} className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
                Upload
            </button>
            {errorMessage && <p className="text-red-500 mt-3">{errorMessage}</p>}
            {presignedURL && (
                <p className="text-green-500 mt-3">
                    File uploaded successfully! <a href={presignedURL} className="text-blue-500">Download</a>
                </p>
            )}
            <ToastContainer />
        </div>
    );
};

export default FileUploader;
