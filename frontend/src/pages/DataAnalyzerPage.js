import { useState } from "react";
import axios from "axios";

const DataAnalyzerPage = () => {
  const [file, setFile] = useState(null);
  const [question, setQuestion] = useState("");
  const [response, setResponse] = useState("");
  const [isLoading, setIsLoading] = useState(false);

  const handleFileChange = (e) => {
    setFile(e.target.files[0]);
  };

  const handleUpload = async () => {
    if (!file || !question) {
      alert("Please upload a file and enter a question.");
      return;
    }

    const formData = new FormData();
    formData.append("file", file);
    formData.append("question", question);

    setIsLoading(true);
    try {
      const res = await axios.post('http://localhost:8080/upload', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });
      setResponse(res.data.answer);
    } catch (error) {
      console.error('Error uploading file:', error);
      setResponse("An error occurred while processing your request.");
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="container mx-auto mt-10 max-w-2xl">
      <h1 className="text-3xl font-bold mb-6 text-center">Data Analyzer</h1>
      <div className="bg-white shadow-md rounded-lg p-6">
        <div className="mb-4">
          <label className="block text-gray-700 font-bold mb-2">Upload Dataset</label>
          <input 
            type="file" 
            onChange={handleFileChange} 
            className="w-full p-2 border rounded"
          />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700 font-bold mb-2">Ask a Question</label>
          <input
            type="text"
            value={question}
            onChange={(e) => setQuestion(e.target.value)}
            placeholder="What insights do you want from this dataset?"
            className="w-full p-2 border rounded"
          />
        </div>
        <button 
          onClick={handleUpload} 
          disabled={isLoading}
          className="w-full bg-blue-500 text-white p-2 rounded hover:bg-blue-600 transition duration-300 disabled:opacity-50"
        >
          {isLoading ? 'Processing...' : 'Analyze Data'}
        </button>
        
        {response && (
          <div className="mt-6 p-4 bg-gray-100 rounded">
            <h2 className="text-xl font-semibold mb-2">Analysis Result</h2>
            <p>{response}</p>
          </div>
        )}
      </div>
    </div>
  );
};

export default DataAnalyzerPage;
