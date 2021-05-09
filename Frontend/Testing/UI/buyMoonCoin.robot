*** Settings ***
Library    Selenium2Library
Library    DateTime
Library    String

*** Variables ***
${browser}      chrome
${url}          http://localhost:3000/buy
${thbt_input_field}     //*[@id="thbt-input"]
${buybtn}      //*[@id="buy-btn"]
${homebtn}    //*[@id="back-btn"]
${balancelabel}     //*[@id="balance-label"]
${historybtn}       //*[@id="root"]/div/div[1]/div/nav/ul/li[2]

*** Keywords ***
Open web browser
    Open browser    ${url}      ${browser}

Input THBT      
    [Arguments]     ${thbt}
    Input text      ${thbt_input_field}         ${thbt}

*** Test Cases ***
Test User must have only 100 THBT
    Open web browser  
    Wait Until Page Contains Element    ${thbt_input_field}      15s
    Element Should Contain      ${balancelabel}         100
    Close browser


Test Buy moon coin
    Open web browser  
    Wait Until Page Contains Element    ${thbt_input_field}      15s
    Input THBT      100
    sleep   5s
    Click Element   ${buybtn}
    Wait Until Page Contains Element     ${homebtn}      15s
    Close browser

Test Buy moon coin and turn back to screen
    Open web browser  
    Wait Until Page Contains Element    ${thbt_input_field}      15s
    Input THBT      100
    sleep   5s
    Click Element   ${buybtn}
    Wait Until Page Contains Element     ${homebtn}      15s
    Click Element   ${homebtn} 
    Wait Until Page Contains Element    ${thbt_input_field}      15s
    Close browser

Test Buy moon coin and check on history must be exist
    Open web browser  
    Wait Until Page Contains Element    ${thbt_input_field}      15s
    Input THBT      100
    sleep   5s
    ${date} =	Get Current Date
    Click Element   ${buybtn}
    Wait Until Page Contains Element     ${homebtn}      15s
    Click Element   ${homebtn} 
    Wait Until Page Contains Element    ${thbt_input_field}      15s
    ${date_to_check}        Get Substring   ${date}     0       -7

    Click Element   ${historybtn}
    Wait Until Page Contains            ${date_to_check}

    Close browser