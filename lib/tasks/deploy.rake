namespace :deploy do
  desc "deploy to staging"
  task staging: :environment do
    sh "git push staging master"
  end

  desc "deploy to production"
  task production: :environment do
    sh "git push production master"
  end
end
