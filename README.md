<a name="readme-top"></a>
  
[![GNU License](https://img.shields.io/badge/license-GNU-brightgreen?style=for-the-badge&logo=gnu)](https://github.com/SomnathDas/whatsapp-botto-go/blob/main/LICENSE)
[![Contributers](https://img.shields.io/github/contributors/SomnathDas/whatsapp-botto-go?style=for-the-badge)](https://github.com/SomnathDas/whatsapp-botto-go)

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/SomnathDas/errand">
    <img src="https://www.svgrepo.com/show/452214/go.svg" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">Whatsapp Botto Go</h3>

  <p align="center">
   A basic automation program written in go-lang using whatsmeow library along with open-ai's CHAT-GPT and ElevenLabs AI TTS.
    <br />
    <br />
    <a href="https://github.com/SomnathDas/whatsapp-botto-go/issues">Report Bug</a>
    ·
    <a href="https://github.com/SomnathDas/whatsapp-botto-go/issues">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

![Meet Valarie](https://github.com/SomnathDas/whatsapp-botto-go/blob/main/welcome/first.png)
![Whatsapp-Botto-Go](https://github.com/SomnathDas/whatsapp-botto-go/blob/main/welcome/second.png)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

* ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
* ![ChatGPT](https://img.shields.io/badge/chatGPT-74aa9c?style=for-the-badge&logo=openai&logoColor=white)
* ![WhatsApp](https://img.shields.io/badge/WhatsApp-25D366?style=for-the-badge&logo=whatsapp&logoColor=white)

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Getting Started

These instructions are mainly written with Linux based operating system in mind.
However these requirements shouldn't differ a lot. Feel free to ask any query.

### Prerequisites

* Go Lang [Read Installation Guide](https://go.dev/doc/install)
* Verify the installation by typing the following command in your console:
* ``go version``
* Install Ffmpeg: [Download Page](https://ffmpeg.org/download.html)
* Install libopus-dev [Download Page](https://opus-codec.org/downloads/)

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/SomnathDas/whatsapp-botto-go.git
   ```
2. Install required GO packages by typing following commands
  ```sh
  cd whatsapp-botto-go
  go get
  ```
3.  Create `.env` file in root dir of the project and enter the following in `.env` file
   ```.env
    OPEN_AI_CHATGPT_API_KEY=""
    ELEVEN_LABS_TTS_API_KEY=""
    ELEVEN_LABS_VOICE_ID=""
    ELEVEN_LABS_MODEL_ID="eleven_monolingual_v1"
    WHATSAPP_NUMBER=""
    AUDIO_FOLDER_ABSOLUTE_PATH=""
   ```
   
   * OPEN_AI_CHATGPT_API_KEY : [OpenAI Website -> Personal -> View API Keys](https://platform.openai.com)
   * ELEVEN_LABS_TTS_API_KEY : [Elevenlabs Website -> Profile -> API Key](https://beta.elevenlabs.io/speech-synthesis)
   * ELEVEN_LABS_VOICE_ID : [Pick your voice model](https://api.elevenlabs.io/v1/voices)
   * WHATSAPP_NUMBER : Example:= 9196XXXXXXXX for India, where 91 is the country code
   * AUDIO_FOLDER_ABSOLUTE_PATH : For example in Linux := ```"/home/somnath/whatsapp-botto-go/audio/"```
   * MY PREFFERED VOICE MODEL "BELLA" : ELEVEN_LABS_VOICE_ID="EXAVITQu4vr4xnSDxMaL"
   
 4. Run the app by typing the following command in root directory of this project
    ```sh
    go run *.go
    ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->
## License

Distributed under the GNU Affero General Public License v3.0 License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- CONTACT -->
## Contact

Somnath Das - [@instagram_handle](https://instagram.com/samurai3247)

Project Link: [https://github.com/SomnathDas/whatsapp-botto-go](https://github.com/SomnathDas/whatsapp-botto-go)

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- ACKNOWLEDGMENTS -->
## Acknowledgments
I whole heartedly thank these amazing men/women for their time and effort.

* [Based on Whatsmeow automation library](https://pkg.go.dev/go.mau.fi/whatsmeow)
* [Using go-openai unofficial library made by sashabaranov](https://pkg.go.dev/github.com/sashabaranov/go-openai)
* [Using eleven-labs unofficial library made by taigrr](https://github.com/taigrr/elevenlabs)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
