import { Outlet, Link } from "react-router-dom";
import "./layout.css"


const Layout = () => {
    return (
        <>
        
            <nav>
                
                <ul>
                    <li>
                        <Link to="/">Home</Link>

                    </li>
                    <li><Link to="/tabular">Tabular</Link></li>
                    <li><Link to="/bar">Barchart</Link></li>
                    <li><Link to="/piegraph">Piechart</Link></li>
                    <li>
                        <Link to="/dashboard">Dashboard</Link>
                    </li>
                    
                </ul>
            </nav>

            <Outlet />
          
        </>
    )
};

export default Layout;