## User Verification

I originally planned to require users to verify their email addresses. In the past, I implemented this using an automated Gmail account, verification token, and a redirect route. However, at some point in 2022, Google changed their Gmail API requirements, which completely broke my registration process.

I want this app to be mostly "set it and forget it", and I don't to deal with anything like that again.  I also still want to prevent the unverified accounts from blowing up my app. For this reason, I've decided to implement a process where users are notified to email `portfolioinstruments@gmail.com` to request manual verification by me.

Since I have an endpoint set up that allows me to easily verify a user based on their email and user ID, I believe this approach will be relatively straightforward and require minimal maintenance.  I think this is an okay solution since I don't expect many users to use this app as it is mostly for me. 