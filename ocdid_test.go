package main

var sampleResponse = `
[
  {
    "district-divisions": [
      {
        "ocd-id": "ocd-division/country:us/state:ma/place:east_longmeadow",
        "voter-registration-authority-level": "municipal",
        "election-authority-level": "municipal",
        "voting-methods": [
          {
            "primary": false,
            "type": "early-voting",
            "excuse-required": false
          },
          {
            "primary": true,
            "instructions": {
              "voting-id": "You may be asked to show identification if: you are voting for the first-time in Massachusetts, you are an inactive voter, you are casting a provisional or challenged ballot, or if the poll worker has a reasonable suspicion that leads them to request identification. Acceptable forms include
(must show your name and address): a MA driver's license or MA-issued ID card; recent utility bill; rent receipt; signed lease; a copy of a voter registration affidavit; or any other printed identification which contains the voter's name and address.\n\nVoters without ID: If you're a first-time voter who is unable to present ID when you check in, you may vote a provisional ballot and return with acceptable ID by close of polls. If you're asked for ID for any other reason, and are not able to present ID in such a situation, you must still be permitted to vote; however, your ballot must be challenged. Your ballot will be cast normally, and will only be re-examined in the case of a recount, court order, or audit."
            },
            "type": "in-person",
            "excuse-required": false
          },
          {
            "primary": false,
            "type": "by-mail",
            "excuse-required": true,
            "ballot-request-deadline-received": "2019-06-03T00:00:00Z",
            "acceptable-forms": [
              {
                "name": "ma_absentee"
              }
            ]
          }
        ],
        "voter-registration-methods": [
          {
            "instructions": {
              "registration": "You should know: you need a Massachusetts ID to use Massachusetts's online voter registration system. If you don't have a Massachusetts-issued ID, you can still register to vote by mail."
            },
            "type": "online",
            "supports-iframe": true,
            "deadline-online": "2019-05-15T00:00:00Z",
            "url": "https://www.sec.state.ma.us/OVR/Pages/MinRequirements.aspx?RMVId=True"
          },
          {
            "deadline-postmarked": "2019-05-15T00:00:00Z",
            "instructions": {
              "idnumber": "Federal law requires that you provide your driver's license number to register to vote. If you do not have a current and valid Massachusetts driver's license, you must provide the last four digits of your Social Security number. If you have neither, you must write \"NONE\" in the box.",
              "signature": "To register in Massachusetts you must: \nbe a citizen of the United States \nbe a resident of Massachusetts \nbe at least 16 years old (must be 18 years old to vote on Election Day)\nnot have been convicted of corrupt practices in respect to elections \nnot be under guardianship with respect to voting \nnot be currently incarcerated for a felony conviction"
            },
            "type": "by-mail",
            "acceptable-forms": [
              {
                "name": "nvrf"
              }
            ]
          }
        ],
        "primary-voting-method-source": "state"
      }
    ],
    "website": "http://www.sec.state.ma.us/wheredoivotema/bal/MyElectionInfo.aspx",
    "polling-place-url": "https://www.sec.state.ma.us/wheredoivotema//bal/myelectioninfo.aspx",
    "date": "2019-06-04T00:00:00Z",
    "population": 15720,
    "polling-place-url-shortened": "https://tvote.org/2v2xAee",
    "description": "East Longmeadow Town Election",
    "id": "5cc0cf70-c9e4-4311-945f-48cde0058291"
  },
  {
    "district-divisions": [
      {
        "ocd-id": "ocd-division/country:us/state:ma/place:peru",
        "voter-registration-authority-level": "municipal",
        "election-authority-level": "municipal",
        "voting-methods": [
          {
            "primary": false,
            "type": "early-voting",
            "excuse-required": false
          },
          {
            "primary": true,
            "instructions": {
              "voting-id": "You may be asked to show identification if: you are voting for the first-time in Massachusetts, you are an inactive voter, you are casting a provisional or challenged ballot, or if the poll worker has a reasonable suspicion that leads them to request identification. Acceptable forms include
(must show your name and address): a MA driver's license or MA-issued ID card; recent utility bill; rent receipt; signed lease; a copy of a voter registration affidavit; or any other printed identification which contains the voter's name and address.\n\nVoters without ID: If you're a first-time voter who is unable to present ID when you check in, you may vote a provisional ballot and return with acceptable ID by close of polls. If you're asked for ID for any other reason, and are not able to present ID in such a situation, you must still be permitted to vote; however, your ballot must be challenged. Your ballot will be cast normally, and will only be re-examined in the case of a recount, court order, or audit."
            },
            "type": "in-person",
            "excuse-required": false
          },
          {
            "primary": false,
            "type": "by-mail",
            "excuse-required": true,
            "ballot-request-deadline-received": "2019-06-07T00:00:00Z",
            "acceptable-forms": [
              {
                "name": "ma_absentee"
              }
            ]
          }
        ],
        "voter-registration-methods": [
          {
            "deadline-postmarked": "2019-05-18T00:00:00Z",
            "instructions": {
              "idnumber": "Federal law requires that you provide your driver's license number to register to vote. If you do not have a current and valid Massachusetts driver's license, you must provide the last four digits of your Social Security number. If you have neither, you must write \"NONE\" in the box.",
              "signature": "To register in Massachusetts you must: \nbe a citizen of the United States \nbe a resident of Massachusetts \nbe at least 16 years old (must be 18 years old to vote on Election Day)\nnot have been convicted of corrupt practices in respect to elections \nnot be under guardianship with respect to voting \nnot be currently incarcerated for a felony conviction"
            },
            "type": "by-mail",
            "acceptable-forms": [
              {
                "name": "nvrf"
              }
            ]
          },
          {
            "instructions": {
              "registration": "You should know: you need a Massachusetts ID to use Massachusetts's online voter registration system. If you don't have a Massachusetts-issued ID, you can still register to vote by mail."
            },
            "type": "online",
            "supports-iframe": true,
            "deadline-online": "2019-05-18T00:00:00Z",
            "url": "https://www.sec.state.ma.us/OVR/Pages/MinRequirements.aspx?RMVId=True"
          }
        ],
        "primary-voting-method-source": "state"
      }
    ],
    "polling-place-url": "https://www.sec.state.ma.us/wheredoivotema//bal/myelectioninfo.aspx",
    "date": "2019-06-08T00:00:00Z",
    "polling-place-url-shortened": "https://tvote.org/2v2xAee",
    "description": "Peru Town Election",
    "id": "5cc0e4a4-8317-4f13-94e1-bdfd153c3af9"
  }
]`
