*** Settings ***
Library    Selenium2Library

*** Variables ***
${browser}      chrome
${url}          http://localhost:3000/buy
${thbt_input_field}     //*[@id="thbt-input"]
${buybtn}      //*[@id="buy-btn"]
${homebtn}    //*[@id="back-btn"]
${balancelabel}     //*[@id="balance-label"]


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
    Click Element   ${buybtn}
    Wait Until Page Contains Element     ${homebtn}      15s
    Click Element   ${homebtn} 
    Wait Until Page Contains Element    ${thbt_input_field}      15s
    Close browser