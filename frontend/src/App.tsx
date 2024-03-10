import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import LandingPage from './pages/LandingPage';
import Error404 from './pages/Error404';
import Register from './pages/Register'
import Dashboard from './pages/Dashboard';
import './App.css';

function App() {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<LandingPage />} />
                <Route path="/register" element={<Register />} />
                <Route path="/dashboard" element={<Dashboard />} />
                {/* Add routes for creating transactions and budgets */}
                {/*<Route path="/create-transaction" element={<CreateTransaction />} />*/}
                {/*<Route path="/create-budget" element={<CreateBudget />} />*/}

                <Route path="*" element={<Error404 />} />
            </Routes>
        </Router>
    );
}

export default App;
