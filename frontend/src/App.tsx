import { useState } from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import LandingPage from './pages/LandingPage';
import Error404 from './pages/Error404';
import './App.css';

function App() {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<LandingPage />} />
                <Route path="*" element={<Error404 />} />
            </Routes>
        </Router>
    );
}

export default App;
