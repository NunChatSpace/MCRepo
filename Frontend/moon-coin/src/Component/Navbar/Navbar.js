import React from 'react';
import { Link } from 'react-router-dom';
import {SidebarData} from "./SidebarData";
import { IconContext } from 'react-icons';
import './Navbar.css';

function Navbar(props) {
    return (
        <div>
            <IconContext.Provider value={{ color: '#fff' }}>
            <nav className='nav-menu active'>
                <ul className='nav-menu-items'>
                    {SidebarData.map((item, index) => {
                        return (
                            <li key={index} className={item.cName}>
                                <Link to={item.path}>
                                    {item.icon}
                                    <span>{item.title}</span>
                                </Link>
                            </li>
                        )
                    })}
                </ul>
                <label style={{color:'white', position:'absolute', bottom:0, padding:5}}>
                    {props.username}
                </label>
            </nav>
            </IconContext.Provider>
        </div>

    )
}

export default Navbar;