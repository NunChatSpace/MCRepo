*** Settings ***
Library    Selenium2Library

*** Variables ***
${browser}      chrome
${url}          http://localhost:3000/buy
${thbt_input_field}     //*[@id="thbt-input"]
${buybtn}      //*[@id="buy-btn"]
${homebtn}    //*[@id="back-btn"]
${balancelabel}     //*[@id="balance-label"]
${historybtn}       //*[@id="root"]/div/div[1]/div/nav/ul/li[2]
${moonleftlabel}     //*[@id="moon-left-label"]

*** Keywords ***
Open web browser
    Open browser    ${url}      ${browser}


To history view
    Click Element   ${historybtn}

*** Test Cases ***
Test to history view
    Open web browser  
    Wait Until Page Contains Element    ${thbt_input_field}      15s
    Element Should Contain      ${balancelabel}         100

    To history view
    Wait Until Page Contains Element    ${moonleftlabel}      15s

    Close browser
    