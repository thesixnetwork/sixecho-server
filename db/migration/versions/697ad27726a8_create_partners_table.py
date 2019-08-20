"""create partners table

Revision ID: 697ad27726a8
Revises: c49e3cfa0654
Create Date: 2019-08-04 17:11:00.055503

"""
import sqlalchemy as sa
from alembic import op
from sqlalchemy.sql import func

# revision identifiers, used by Alembic.
revision = '697ad27726a8'
down_revision = 'c49e3cfa0654'
branch_labels = None
depends_on = None


def upgrade():
    op.create_table(
        'partners',
        sa.Column('api_key', sa.String(255), primary_key=True),
        sa.Column('api_secret', sa.String(255), nullable=False),
        sa.Column('name', sa.String(255), nullable=False),
        sa.Column('created_at',
                  sa.DateTime(timezone=True),
                  nullable=False,
                  server_default=func.now()),
        sa.Column('updated_at',
                  sa.DateTime(timezone=True),
                  nullable=False,
                  server_default=func.now()),
    )


def downgrade():
    op.drop_table('partners')
