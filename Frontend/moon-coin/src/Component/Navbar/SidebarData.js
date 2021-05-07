import React from 'react'
import * as AIIcons from "react-icons/ai"
import * as BIIcons from "react-icons/bi"

export const SidebarData = [
    {
        title: 'Buy',
        path:'/buy',
        icon: <BIIcons.BiTransferAlt/>,
        cName: 'nav-text'
    },
    {
        title: 'History',
        path:'/history',
        icon: <AIIcons.AiOutlineHistory/>,
        cName: 'nav-text'
    }
];