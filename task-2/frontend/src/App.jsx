import { useState, useEffect } from "react";
import axios from "axios";
import "./App.css";

function App() {
  const [name, setName] = useState("");
  const [mobile, setMobile] = useState("");
  const [submittedData, setSubmittedData] = useState([]);

  useEffect(() => {
    fetchUsers();
  }, []);

  const fetchUsers = async () => {
    try {
      const res = await axios.get("http://127.0.0.1:49795/users");
      setSubmittedData(res.data);
    } catch (error) {
      console.error("Error fetching users:", error);
    }
  };

  const handleSubmit = async () => {
    if (!name || !mobile) return alert("Please fill all fields");

    try {
      await axios.post("http://127.0.0.1:49795/add", {
        name: name,
        password: mobile,
      });
      alert("User added successfully");
      setName("");
      setMobile("");
      fetchUsers();
    } catch (error) {
      console.error("Error adding user:", error);
      alert("Failed to add user");
    }
  };

  return (
    <div className="container">
      {/* Left Side */}
      <div className="left">
        <h2>Submitted Data</h2>
        {submittedData && submittedData.length === 0 ? (
          <p>No data submitted yet</p>
        ) : (
          <div className="data-list">
            {submittedData && submittedData.map((data, index) => (
              <div key={index} className="data-card">
                <div>
                  <strong>Name:</strong> {data.name}
                </div>
                <div>
                  <strong>Mobile:</strong> {data.password}
                </div>
              </div>
            ))}
          </div>
        )}
      </div>

      {/* Right Side */}
      <div className="right">
        <div className="form">
          <h2>Enter Details</h2>
          <label>Name</label>
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="Enter your name"
          />

          <label>Mobile</label>
          <input
            type="tel"
            value={mobile}
            onChange={(e) => setMobile(e.target.value)}
            placeholder="Enter your mobile"
          />

          <button onClick={handleSubmit}>Submit</button>
        </div>
      </div>
    </div>
  );
}

export default App;
