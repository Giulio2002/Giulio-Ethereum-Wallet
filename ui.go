package main

import (
	"github.com/andlabs/ui"
	"crypto/sha256"
	"io"
	"os"
	"fmt"
	"log"
)

func index(){
	err := ui.Main(func() {
		txButton := ui.NewButton("tx");
		personalButton := ui.NewButton("personal");
		box := ui.NewVerticalBox();
		box.Append(txButton,false);
		box.Append(personalButton,false);
		window := ui.NewWindow("Rebuffo Wallet", 200, 100,false);
		window.SetMargined(true);
		window.SetChild(box);
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		txButton.OnClicked(func(*ui.Button){
			window.Destroy();
			go txPage();
		})
		
		personalButton.OnClicked(func(*ui.Button){
			window.Destroy();
			go personalPage();
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}

func txPage(){
	err := ui.Main(func() {
		back := ui.NewButton("Back");
		from := ui.NewEntry();
		to := ui.NewEntry();
		amount := ui.NewEntry();
		password := ui.NewEntry();
		/*quick setup*/
		info1 := ui.NewLabel("from:");
		info2 := ui.NewLabel("to:");
		info3 := ui.NewLabel("amount:");
		info4 := ui.NewLabel("password:");
		commitButton := ui.NewButton("commit");
		box := ui.NewVerticalBox();
		box.Append(back,false);
		box.Append(info1,false);
		box.Append(from,false);
		box.Append(info2,false);
		box.Append(to,false);
		box.Append(info3,false);
		box.Append(amount,false);
		box.Append(info4,false);
		box.Append(password,false);
		box.Append(commitButton,false);
		window := ui.NewWindow("Rebuffo Wallet", 500, 100,false);
		window.SetMargined(true);
		window.SetChild(box);
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit();
			return true;
		})
		commitButton.OnClicked(func(*ui.Button){
			hasher := sha256.New();
			file,err := os.Open(password.Text());
			if err != nil && password.Text() != ""{
				log.Fatal(err)
			}
			defer file.Close();
			io.Copy(hasher,file);
			sum := string(hasher.Sum(nil))
			fmt.Println(sum);
			PersonalUnlockAccount(from.Text(),sum,50);
			tx := TransactionObject{};
			tx.From = from.Text();
			tx.To = to.Text();
			tx.Value = amount.Text();
			tx.Gas = "31000";
			EthSendTransaction(&tx);
		})

		back.OnClicked(func(*ui.Button){
			window.Destroy();
			go index();
		})		
		window.Show()
	})
	if err != nil {
		panic(err)
	}	
}

func personalPage(){
	err := ui.Main(func() {

		back := ui.NewButton("Back");
		password := ui.NewEntry();
		info := ui.NewLabel("Create new account:");
		created := ui.NewEntry();
		created.SetReadOnly(true);
		commitButton := ui.NewButton("commit");
		box := ui.NewVerticalBox();
		box.Append(back,false);
		box.Append(info,false);
		box.Append(password,false);
		box.Append(commitButton,false);
		box.Append(created,false);
		window := ui.NewWindow("Rebuffo Wallet", 500, 100,false);
		window.SetMargined(true);
		window.SetChild(box);
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit();
			return true;
		})

		commitButton.OnClicked(func(*ui.Button){
			hasher := sha256.New();
			file,err := os.Open(password.Text());
			if err != nil{
				log.Fatal(err)
			}
			defer file.Close();
			io.Copy(hasher,file);
			sum := string(hasher.Sum(nil))
			fmt.Println(sum);
			created.SetText("new address at " + PersonalNewAccount(sum));
		})

		back.OnClicked(func(*ui.Button){
			window.Destroy();
			go index();
		})		
		window.Show()
	})
	if err != nil {
		panic(err)
	}	
}