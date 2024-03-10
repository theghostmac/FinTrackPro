import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import './LandingPage.css';

interface User {
    UserID: string;
    userName: string;
    email: string;
    Transactions: object,
}

const Dashboard = () => {
    const [userData, setUserData] = useState<User | null>(null);

    useEffect(() => {
        const fetchProfile = () => {
            const token = localStorage.getItem('token');
            fetch('http://localhost:9020/profile', {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            })
                .then(response => {
                    if (!response.ok) {
                        throw new Error('Error fetching profile');
                    }
                    return response.json();
                })
                .then(data => {
                    setUserData(data);
                })
                .catch(error => {
                    console.error('There was an error fetching the profile:', error);
                });
        };

        fetchProfile();
    }, []);

    return (
        <div>
            <h1>Dashboard</h1>
            <p>Welcome back, {userData?.userName}! Here's your financial overview.</p>
            {/* Display transactions and other user data */}
            <div className="actions">
                <Link to="/create-transaction" className="dashboard-button">New Transaction</Link>
                <Link to="/create-budget" className="dashboard-button">New Budget</Link>
            </div>
        </div>
    );
}

export default Dashboard;