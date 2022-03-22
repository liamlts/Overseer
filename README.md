# Overseer
<h4>Dependencies</h4>
  <li><i>
  Uncomplicated firewall.(UFW)
   </li>
  <li>
  Password authentication disabled in your sshd_config file.
  </li></i>
  <h4>About</h4>
  <p>
  This program scans your log file and checks for malicious hosts whether it be a scanner such as shodan or a bruteforce/dictionary attack. Then usuing
  KeyCDNs geolocation API and UFW displays info on the malicious hosts and blocks them from further access to your server.
  </p>
  <h4>Running Overseer</h4>

  
  ```
  go build overseer.go
  sudo ./overseer
  ```

Please note this is far from a finished project.
