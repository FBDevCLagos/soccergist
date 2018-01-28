# HOW TO CREATE A FACEBOOK APP
_- Jeremiah Heir_
The second step to take when creating your bot is to create a Facebook app. The Facebook app you'll create will be similar to apps used for every other Facebook product offerings such as [Facebook Login](https://developers.facebook.com/docs/facebook-login), [Account kit](https://developers.facebook.com/docs/accountkit), [Facebook Sharing](https://developers.facebook.com/docs/sharing/overview) etc. The significant difference is you will add and use the Messenger product addon to your app.

With the Messenger product addon, your Facebook app will serve as the connection between your Facebook page and the hosted backend for your bot. The addon will allow you specify necessary information and configurations such as 
* [webhook](https://developers.facebook.com/docs/messenger-platform/webhook#setup) for your hosted backend i.e the url where your bot is hosted
* [type of events](https://developers.facebook.com/docs/messenger-platform/webhook#setup) your bot is interested in
* Facebook page to listen on for these events

You can find more information on webhooks in the [next article]() and in [Part 2](part2) on messenger events. For now, we will focus on creating an app with Messenger setup.

You'll need a Facebook developer account to create an app. To confirm if you have an account, visit [https://developers.facebook.com](https://developers.facebook.com) and log in. If you don't have one, upgrade your personal Facebook account to a developer account by visiting [this link](https://developers.facebook.com/async/onboarding/dialog/)

Now let's create our app.

You can follow the official documentation [here](https://developers.facebook.com/docs/apps/register/) but we'll re-echo it for redundancy.

1. [Login to Facebook](https://www.facebook.com/login.php)
2. [Developer Account](https://developers.facebook.com)
3. [Visit the Facebook developer section to create new Facebook app](https://developers.facebook.com/apps)
    * Choose **Apps** in the header navigation and select _Add a New App_ or use [this link](https://developers.facebook.com/apps/async/create/platform-setup/dialog/).

    * Enter a display name and contact email address for your app

        ![enter app details screenshot](https://user-images.githubusercontent.com/11221027/29972391-e16d732e-8f23-11e7-8d95-3de7dd4ec056.png)

    * Click the create app id button and follow through with the security check

    Your app is now created. Completed the following steps to add the Messenger product addon to your app

    * From the left menu, choose the **+Add Product** option
    * From the product listings find the **Messenger product** and click the **set up** button. 

        ![setup messenger screenshot](https://user-images.githubusercontent.com/11221027/29972467-14945ccc-8f24-11e7-82f7-ae3db523d2a4.png)

    This adds the messenger product to your app and lists it under the products section on the right menu. 

    And that's all needed to create a Facebook app with Messenger setup. Next todo is [setting up your messenger webhook](#).

